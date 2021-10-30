package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/manerakai/simplytranslate_go/engines"
)

func translate(w http.ResponseWriter, r *http.Request) {

	text := r.FormValue("text")
	from := r.FormValue("from")
	to := r.FormValue("to")

	if text == "" {
		writeError(w, 400, "'text' field is missing")
	}
	if from == "" {
		writeError(w, 400, "'from' field is missing")
	}
	if to == "" {
		writeError(w, 400, "'to' field is missing")
	}

	fmt.Fprintf(w, engines.GoogleTranslateEngine.Translate(text, from, to))

}

func getSupportedLanguages(w http.ResponseWriter, r *http.Request) {
	engine := r.FormValue("engine")

	request, error := json.Marshal(engines.GetSupportedLanguages(engine))

	if error != nil {
		fmt.Println(error)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(request))

}

func tts(w http.ResponseWriter, r *http.Request) {

	text := r.FormValue("text")
	lang := r.FormValue("lang")

	if text == "" {
		writeError(w, 400, "'text' field is missing")
	}
	if lang == "" {
		writeError(w, 400, "'from' field is missing")
	}
	if lang == "" || text == "" {
		w.Header().Set("Content-Type", "application/html")
		fmt.Fprint(w, "")
	} else {
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(engines.GoogleTranslateEngine.Tts(text, lang))
	}
}
