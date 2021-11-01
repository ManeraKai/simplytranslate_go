package main

import (
	"fmt"
	"html/template"
	"net/http"
	"simplytranslate_go/engines"
)

type indexDataStruct struct {
	LangList map[string]string
}

func index(w http.ResponseWriter, r *http.Request) {

	var langList map[string]string = make(map[string]string)

	for k, v := range engines.GetSupportedLanguages("Google") {
		langList[k] = v
	}

	indexData := indexDataStruct{
		langList,
	}

	http.FileServer(http.Dir("/"))
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
	fmt.Println("Running on http://localhost:8090/")
	http.ListenAndServe(":8090", nil)
}
