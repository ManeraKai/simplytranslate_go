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

type libreTranslateEngineStruct struct{}

var LibreTranslateEngine = func() *libreTranslateEngineStruct {
	return &libreTranslateEngineStruct{}
}()

var libreSupportedLanguages map[string]string

func (self libreTranslateEngineStruct) Name() string {
	return "libre"
}

func (self libreTranslateEngineStruct) GetSupportedLanguages() map[string]string {
	if libreSupportedLanguages != nil {
		return libreSupportedLanguages
	}

	resp, err := http.Get("https://almaleehserver.asuscomm.com:451/languages")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bodyBytes, error := ioutil.ReadAll(resp.Body)

	if error != nil {
		log.Fatal(error)
	}

	bodyString := string(bodyBytes)

	// This is for testing to not abuse libretranslate's api
	// bodyString := "[{\"code\":\"en\",\"name\":\"English\"},{\"code\":\"ar\",\"name\":\"Arabic\"},{\"code\":\"zh\",\"name\":\"Chinese\"},{\"code\":\"nl\",\"name\":\"Dutch\"},{\"code\":\"fi\",\"name\":\"Finnish\"},{\"code\":\"fr\",\"name\":\"French\"},{\"code\":\"de\",\"name\":\"German\"},{\"code\":\"hi\",\"name\":\"Hindi\"},{\"code\":\"hu\",\"name\":\"Hungarian\"},{\"code\":\"id\",\"name\":\"Indonesian\"},{\"code\":\"ga\",\"name\":\"Irish\"},{\"code\":\"it\",\"name\":\"Italian\"},{\"code\":\"ja\",\"name\":\"Japanese\"},{\"code\":\"ko\",\"name\":\"Korean\"},{\"code\":\"pl\",\"name\":\"Polish\"},{\"code\":\"pt\",\"name\":\"Portuguese\"},{\"code\":\"ru\",\"name\":\"Russian\"},{\"code\":\"es\",\"name\":\"Spanish\"},{\"code\":\"sv\",\"name\":\"Swedish\"},{\"code\":\"tr\",\"name\":\"Turkish\"},{\"code\":\"uk\",\"name\":\"Ukranian\"},{\"code\":\"vi\",\"name\":\"Vietnamese\"}]"

	lengthy, lengthyError := strconv.Atoi(gjson.Get(bodyString, "#").Raw)

	if lengthyError != nil {
		log.Fatal(lengthyError)
	}

	var result map[string]string = make(map[string]string)

	for i := 0; i < lengthy; i++ {
		element := gjson.Get(bodyString, strconv.Itoa(i))
		result[element.Get("code").Str] = element.Get("name").Str
	}

	return result

}

func (self libreTranslateEngineStruct) Translate(text, from, to string) string {

	values := map[string]string{"q": strings.TrimSpace(text), "source": from, "target": to}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, httpError := http.Post("https://almaleehserver.asuscomm.com:451/translate?", "application/json", bytes.NewBuffer(json_data))

	if httpError != nil {
		fmt.Println("httpError:", httpError)
	}

	defer resp.Body.Close()

	bodyBytes, error := ioutil.ReadAll(resp.Body)

	if error != nil {
		log.Fatal(error)
	}

	bodyString := string(bodyBytes)

	// For testing without abusing the api
	// bodyString := "{\"translatedText\": \"¡Hola mundo raro!\"}"
	// bodyString := "{\"error\": \"error\"}"

	if gjson.Get(bodyString, "error").Str != "" {
		return ""
	}

	return gjson.Get(bodyString, "translatedText").Str
}

func (self libreTranslateEngineStruct) DetectLanguage(text string) string {

	values := map[string]string{"q": strings.TrimSpace(text)}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, httpError := http.Post("https://almaleehserver.asuscomm.com:451/detect", "application/json", bytes.NewBuffer(json_data))

	if httpError != nil {
		fmt.Println("httpError:", httpError)
	}

	defer resp.Body.Close()

	bodyBytes, error := ioutil.ReadAll(resp.Body)

	if error != nil {
		log.Fatal(error)
	}

	bodyString := string(bodyBytes)

	// For testing without abusing the api
	// bodyString := "{\"language\": \"en\"}"
	// bodyString := "{\"error\": \"error\"}"

	if gjson.Get(bodyString, "error").Str != "" {
		return ""
	}

	return gjson.Get(bodyString, "0").Get("language").Str
}
