package simplytranslate_engines

import (
	"strings"
)

type utils struct {
}

var Utils = func() *utils {
	return &utils{}
}()

func (self *utils) ToFullName(langCode string, engine Engine) string {
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

func (self *utils) ToLangCode(lang string, engine Engine) string {
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
