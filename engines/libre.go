package engines

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

type libreTranslateEngineStruct struct {
	Name string
}

var LibreTranslateEngine = func() *libreTranslateEngineStruct {
	return &libreTranslateEngineStruct{"libre"}
}()

func (self libreTranslateEngineStruct) GetSupportedLanguages() map[string]string {
	return supportedLangs
}

func (self libreTranslateEngineStruct) Translate(text, from, to string) string {

	values := map[string]string{"q": strings.TrimSpace(text), "source": from, "target": to}

	if apiKey != "" {
		values["api_key"] = apiKey
	}

	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, httpError := http.Post(instance+"/translate?", "application/json", bytes.NewBuffer(json_data))

	if httpError != nil {
		fmt.Println("httpError:", httpError)
	}

	defer resp.Body.Close()

	bodyBytes, error := ioutil.ReadAll(resp.Body)

	if error != nil {
		log.Fatal(error)
	}

	bodyString := string(bodyBytes)

	if gjson.Get(bodyString, "error").Str != "" {
		return ""
	}

	return gjson.Get(bodyString, "translatedText").Str
}

func (self libreTranslateEngineStruct) DetectLanguage(text string) string {

	values := map[string]string{"q": strings.TrimSpace(text)}

	if apiKey != "" {
		values["api_key"] = apiKey
	}

	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, httpError := http.Post(instance+"/detect", "application/json", bytes.NewBuffer(json_data))

	if httpError != nil {
		fmt.Println("httpError:", httpError)
	}

	defer resp.Body.Close()

	bodyBytes, error := ioutil.ReadAll(resp.Body)

	if error != nil {
		log.Fatal(error)
	}

	bodyString := string(bodyBytes)

	if gjson.Get(bodyString, "error").Str != "" {
		return ""
	}

	return gjson.Get(bodyString, "0").Get("language").Str
}

func initSupportedLangs() map[string]string {
	resp, err := http.Get(instance + "/languages")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bodyBytes, error := ioutil.ReadAll(resp.Body)

	if error != nil {
		log.Fatal(error)
	}

	bodyString := string(bodyBytes)

	lengthy, lengthyError := strconv.Atoi(gjson.Get(bodyString, "#").Raw)

	if lengthyError != nil {
		fmt.Println(lengthyError)
	}

	var result map[string]string = make(map[string]string)

	for i := 0; i < lengthy; i++ {
		element := gjson.Get(bodyString, strconv.Itoa(i))
		result[element.Get("code").Str] = element.Get("name").Str
	}
	return result
}

var instance string
var apiKey string
var supportedLangs map[string]string

func init() {
	instance = config.GetString("libre.instance")
	apiKey = config.GetString("libre.apiKey")
	supportedLangs = initSupportedLangs()
}
