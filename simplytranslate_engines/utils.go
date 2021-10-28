package simplytranslate_engines

import (
	"strings"
)

type utils struct {
}

func newUtilsStruct() *utils {
	return &utils{}
}

var Utils = newUtilsStruct()

func Get_engine(engine_name string, engines []struct {
	Name                    string
	Get_supported_languages func() map[string]string
	Detect_language         func(text string)
	Get_tts                 func(text string, language string) string
	Translate               func(text string, to_language string, from_language string) string
}, default_engine struct {
	Name                    string
	Get_supported_languages func() map[string]string
	Detect_language         func(text string)
	Get_tts                 func(text string, language string) string
	Translate               func(text string, to_language string, from_language string) string
}) string {
	for engine := range engines {
		if engines[engine].Name == engine_name {
			return engines[engine].Name
		}
	}
	return default_engine.Name
}

func (self *utils) To_full_name(lang_code string, engine struct {
	Name                    string
	Get_supported_languages func() map[string]string
	Detect_language         func(text string)
	Get_tts                 func(text string, language string) string
	Translate               func(text string, to_language string, from_language string) string
},
) string {
	lang_code = strings.ToLower(lang_code)
	if lang_code == "auto" {
		return "Autodetect"
	}
	for key, value := range engine.Get_supported_languages() {
		if strings.ToLower(value) == lang_code {
			return key
		}
	}
	return ""
}

func (self *utils) To_lang_code(lang string, engine struct {
	Name                    string
	Get_supported_languages func() map[string]string
	Detect_language         func(text string)
	Get_tts                 func(text string, language string) string
	Translate               func(text string, to_language string, from_language string) string
}) string {
	lang = strings.ToLower(lang)

	if lang == "autodetect" || lang == "auto" {
		return "auto"
	}

	for key, value := range engine.Get_supported_languages() {
		if strings.ToLower(key) == lang {
			return value
		}
	}
	return ""
}
