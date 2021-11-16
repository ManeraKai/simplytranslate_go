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
	Output       string
	Engine       string
	From         string
	To           string
	Input        string
}

func index(w http.ResponseWriter, r *http.Request) {

	from := r.FormValue("from_language")
	to := r.FormValue("to_language")
	text := r.FormValue("input")
	switchlangs := r.FormValue("switchlangs")
	engine := r.FormValue("engine")

	if engine != "google" && engine != "libre" {
		engine = "google"
	}

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

	var output string

	if r.Method == http.MethodGet {
		for _, c := range r.Cookies() {
			if c.Name == "from" {
				from = c.Value
			}
			if c.Name == "to" {
				to = c.Value
			}
		}
	} else if r.Method == http.MethodPost {
		fromCookie := &http.Cookie{
			Name:   "from",
			Value:  from,
			MaxAge: 300,
		}
		toCookie := &http.Cookie{
			Name:   "to",
			Value:  to,
			MaxAge: 300,
		}
		http.SetCookie(w, fromCookie)
		http.SetCookie(w, toCookie)
		output = engines.Translate(text, from, to, engine)
	}

	var langListFrom []string = []string{}

	isAutoSelected := ""

	if from == "auto" {
		isAutoSelected = "selected"
	}

	langListFrom = append(langListFrom, fmt.Sprintf("<option %s value=\"auto\">AutoDetect</option>", isAutoSelected))

	var langListTo []string = []string{}

	for k, v := range engines.GetSupportedLanguages(engine) {
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
		output,
		engine,
		from,
		to,
		text,
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
