package simplytranslate_engines

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
