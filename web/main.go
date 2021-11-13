package main

import (
	"fmt"
	"net/http"
	"simplytranslate_go/engines"
	"sort"
	"text/template"
)

type indexDataStruct struct {
	LangListFrom []string
	LangListTo   []string
	Input        string
	Output       string
	From         string
	To           string
}

func index(w http.ResponseWriter, r *http.Request) {

	from := r.FormValue("from_language")
	to := r.FormValue("to_language")
	text := r.FormValue("input")
	switchlangs := r.FormValue("switchlangs")

	if from == "" {
		from = "auto"
	}
	if to == "" {
		to = "en"
	}

	if switchlangs == "true" && from != "auto" {
		tmpFrom := from
		from = to
		to = tmpFrom
	}

	output := engines.Translate(text, from, to, "google")

	var langListFrom []string = []string{}

	isAutoSelected := ""

	if from == "auto" {
		isAutoSelected = "selected"
	}

	langListFrom = append(langListFrom, fmt.Sprintf("<option %s value=\"auto\">AutoDetect</option>", isAutoSelected))

	var langListTo []string = []string{}

	for k, v := range engines.GetSupportedLanguages("Google") {
		isSelectedFrom := ""
		isSelectedTo := ""
		if k == from {
			isSelectedFrom = "selected"
		}
		if k == to {
			isSelectedTo = "selected"
		}

		langListFrom = append(langListFrom, fmt.Sprintf("<option %s value=\"%s\">%s</option>", isSelectedFrom, k, v))
		langListTo = append(langListTo, fmt.Sprintf("<option %s value=\"%s\">%s</option>", isSelectedTo, k, v))
		sort.Strings(langListFrom)
		sort.Strings(langListTo)
	}

	indexData := indexDataStruct{
		langListFrom,
		langListTo,
		text,
		output,
		from,
		to,
	}

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, indexData)
}

func writeError(w http.ResponseWriter, status int, err interface{}) {
	w.WriteHeader(status)
	fmt.Fprintln(w, err)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/translate", translate)
	http.HandleFunc("/api/get_languages", getSupportedLanguages)
	http.HandleFunc("/api/tts", tts)
	fmt.Println("Running on http://localhost:8097/")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8097", nil)
}
