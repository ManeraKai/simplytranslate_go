package main

import (
	"fmt"
	"net/http"
)

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
