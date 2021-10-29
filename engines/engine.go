package engines

import (
	"strings"
)

type Engine interface {
	Name() string
	GetSupportedLanguages() map[string]string
	Translate(text, from, to string) string
}

func GetSupportedLanguages(engine Engine) map[string]string {
	return engine.GetSupportedLanguages()
}

func Translate(text, from, to string, engine Engine) string {
	return engine.Translate(text, from, to)
}

func ToFullName(langCode string, engine Engine) string {
	langCode = strings.ToLower(langCode)
	if langCode == "auto" {
		return "Autodetect"
	}
	for key, value := range engine.GetSupportedLanguages() {
		if strings.ToLower(key) == langCode {
			return value
		}
	}
	return ""
}

func ToLangCode(lang string, engine Engine) string {
	lang = strings.ToLower(lang)

	if lang == "autodetect" || lang == "auto" {
		return "auto"
	}

	for key, value := range engine.GetSupportedLanguages() {
		if strings.ToLower(value) == lang {
			return key
		}
	}
	return ""
}
