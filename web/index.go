package main

import (
	"fmt"
	"net/http"
	"simplytranslate_go/engines"
	"sort"
	"strings"
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
	EngineList   []string
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

		if switchlangs == "switch" {
			tmpFrom := from
			from = to
			to = tmpFrom
		}

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
		if text != "" {
			output = engines.Translate(text, from, to, engine)
		}
	}

	isAutoSelected := ""

	if from == "auto" {
		isAutoSelected = "selected"
	}

	var langListFrom []string = []string{}
	var langListTo []string = []string{}
	langListFrom = append(langListFrom, fmt.Sprintf("<option %s value=\"auto\">AutoDetect</option>", isAutoSelected))

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

	var engineList []string

	if config.GetBool("google.enabled") {
		engineList = append(engineList, "<a href=\"/?engine=google\">Google</a>&nbsp;|")
	}
	if config.GetBool("libre.enabled") {
		engineList = append(engineList, "<a href=\"/?engine=libre\">LibreTranslate</a>&nbsp;|")
	}

	lastIndex := len(engineList) - 1

	engineList[lastIndex] = strings.ReplaceAll(engineList[lastIndex], "&nbsp;|", "")

	indexData := indexDataStruct{
		langListFrom,
		langListTo,
		output,
		engine,
		from,
		to,
		text,
		engineList,
	}

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, indexData)
}
