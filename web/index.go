package main

import (
	"fmt"
	htmlTemplate "html/template"
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

func mapkey(m map[string]string, value string) (key string, ok bool) {
	for k, v := range m {
		if v == value {
			key = k
			ok = true
			return
		}
	}
	return
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
	var langList []string = []string{}
	var langListVals []string = []string{}

	isSelectedFrom := ""
	isSelectedTo := ""

	mapList := engines.GetSupportedLanguages(engine)
	for _, v := range mapList {
		langList = append(langList, v)
		langListVals = append(langListVals, v)
	}

	sort.Strings(langListVals)

	for _, v := range langListVals {
		k, _ := mapkey(mapList, v)

		if k == from {
			isSelectedFrom = "selected"
		} else {
			isSelectedFrom = ""
		}
		langListFrom = append(langListFrom, fmt.Sprintf("<option %s value=\"%s\" >%s</option>", isSelectedFrom, k, v))
		if k == to {
			isSelectedTo = "selected"
		} else {
			isSelectedTo = ""
		}
		langListTo = append(langListTo, fmt.Sprintf("<option %s value=\"%s\" >%s</option>", isSelectedTo, k, v))
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

	text = htmlTemplate.HTMLEscapeString(text)
	output = htmlTemplate.HTMLEscapeString(output)

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
