package engines

import (
	"strings"

	"github.com/spf13/viper"
)

type Engine interface {
	GetSupportedLanguages() map[string]string
	Translate(text, from, to string) string
}

func isEngineEnabled(engine string) bool {
	if conf.GetBool(engine+".enabled") == true {
		return true
	}
	return false
}

func GetSupportedLanguages(engine string) map[string]string {
	if isEngineEnabled(engine) {
		return GetEngine(engine).GetSupportedLanguages()
	}
	return make(map[string]string)
}

func Translate(text, from, to string, engine string) string {
	if isEngineEnabled(engine) {
		return GetEngine(engine).Translate(text, from, to)
	}
	return ""
}

func ToFullName(langCode string, engine string) string {
	if isEngineEnabled(engine) {
		langCode = strings.ToLower(langCode)
		if langCode == "auto" {
			return "Autodetect"
		}
		for key, value := range GetEngine(engine).GetSupportedLanguages() {
			if strings.ToLower(key) == langCode {
				return value
			}
		}
	}
	return ""
}

func ToLangCode(lang string, engine string) string {
	if isEngineEnabled(engine) {

		lang = strings.ToLower(lang)

		if lang == "autodetect" || lang == "auto" {
			return "auto"
		}

		for key, value := range GetEngine(engine).GetSupportedLanguages() {
			if strings.ToLower(value) == lang {
				return key
			}
		}
	}
	return ""
}

func GetEngine(engineName string) Engine {

	switch engineName {
	case "google":
		return GoogleTranslateEngine
	case "libre":
		return LibreTranslateEngine
	default:
		return GoogleTranslateEngine
	}

}

var conf *viper.Viper

func init() {
	conf = viper.New()
	conf.SetConfigFile("/etc/simplytranslate_go/web.yaml")
	conf.ReadInConfig()
}
