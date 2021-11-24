package main

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

func writeError(w http.ResponseWriter, status int, err interface{}) {
	w.WriteHeader(status)
	fmt.Fprintln(w, err)
}

var config *viper.Viper

func main() {
	config = viper.New()
	config.SetConfigFile("/etc/simplytranslate_go/web.yaml")
	config.ReadInConfig()

	http.HandleFunc("/", index)
	http.HandleFunc("/api/translate", translate)
	http.HandleFunc("/api/get_languages", getSupportedLanguages)
	http.HandleFunc("/api/tts", tts)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	port := config.GetString("network.port")
	cert := config.GetString("network.certificate")
	privKey := config.GetString("network.privateKey")

	if cert != "" && privKey != "" {
		fmt.Println("Running on https://localhost:" + port + "/")
		http.ListenAndServeTLS(":"+port, cert, privKey, nil)
	} else {
		fmt.Println("Running on http://localhost:" + port + "/")
		http.ListenAndServe(":"+port, nil)
	}

}
