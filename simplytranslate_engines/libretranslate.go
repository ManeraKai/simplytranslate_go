package simplytranslate_engines

import (
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

type libreTranslateEngineStruct struct {
	Name                    string
	Get_supported_languages func() map[string]string
	Detect_language         func(text string)
	Get_tts                 func(text string, language string) string
	Translate               func(text string, to_language string, from_language string) string
}

var libreUrl = "https://libretranslate.de"

var supported_languages map[string]string = make(map[string]string)

type LangsStruct struct {
	Lang []map[string]string
}

func newlibreTranslateEngineStruct() *libreTranslateEngineStruct {
	return &libreTranslateEngineStruct{

		"libre",
		func() map[string]string {
			response, _ := http.Get(libreUrl + "/languages")
			bodyBytes, _ := ioutil.ReadAll(response.Body)
			bodyBytesString := string(bodyBytes)

			resultMap := make(map[string]string)

			gjson.ForEachLine(bodyBytesString, func(line gjson.Result) bool {
				line.ForEach(func(_, value gjson.Result) bool {
					resultMap[value.Get("name").Str] = value.Get("code").Str
					return true
				})
				return true
			})

			return resultMap
		},
		func(text string) {
			return
		},
		func(text, language string) string {
			return ""
		},
		func(text, to_language, from_language string) string {
			return ""
		},
	}

}

var LibreTranslateEngine = newlibreTranslateEngineStruct()
