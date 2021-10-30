package engines

import (
	"strings"
)

type Engine interface {
	Name() string
	GetSupportedLanguages() map[string]string
	Translate(text, from, to string) string
}

func GetSupportedLanguages(engine string) map[string]string {
	return GetEngine(engine).GetSupportedLanguages()
}

func Translate(text, from, to string, engine string) string {
	return GetEngine(engine).Translate(text, from, to)
}

func ToFullName(langCode string, engine string) string {
	langCode = strings.ToLower(langCode)
	if langCode == "auto" {
		return "Autodetect"
	}
	for key, value := range GetEngine(engine).GetSupportedLanguages() {
		if strings.ToLower(key) == langCode {
			return value
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

func ToLangCode(lang string, engine string) string {
	lang = strings.ToLower(lang)

	if lang == "autodetect" || lang == "auto" {
		return "auto"
	}

	for key, value := range GetEngine(engine).GetSupportedLanguages() {
		if strings.ToLower(value) == lang {
			return key
		}
	}
	return ""
}
