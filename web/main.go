package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/manerakai/simplytranslate_go/engines"
)

func index(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("index.html")
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(file)
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Fprintf(w, string(b))
}

func writeError(w http.ResponseWriter, status int, err interface{}) {
	w.WriteHeader(status)
	fmt.Fprintln(w, err)
}

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

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/translate", translate)
	http.HandleFunc("/api/get_languages", getSupportedLanguages)
	fmt.Println("Running on http://localhost:3000/")
	http.ListenAndServe(":3000", nil)
}
