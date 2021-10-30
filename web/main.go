package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

	http.FileServer(http.Dir("/"))

	fmt.Fprintf(w, string(b))
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
	fmt.Println("Running on http://localhost:3000/")
	http.ListenAndServe(":3000", nil)
}
