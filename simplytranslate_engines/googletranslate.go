package simplytranslate_engines

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type googleTranslateEngineStruct struct {
	Name                    string
	Get_supported_languages func() map[string]string
	Detect_language         func(text string)
	Get_tts                 func(text string, language string) string
	Translate               func(text string, to_language string, from_language string) string
}

func newgoogleTranslateEngineStruct() *googleTranslateEngineStruct {
	return &googleTranslateEngineStruct{
		"google",
		func() map[string]string {
			return map[string]string{
				"Afrikaans":          "af",
				"Albanian":           "sq",
				"Amharic":            "am",
				"Arabic":             "ar",
				"Armenian":           "hy",
				"Azerbaijani":        "az",
				"Basque":             "eu",
				"Belarusian":         "be",
				"Bengali":            "bn",
				"Bosnian":            "bs",
				"Bulgarian":          "bg",
				"Catalan":            "ca",
				"Cebuano":            "ceb",
				"Chichewa":           "ny",
				"Chinese":            "zh-CN",
				"Corsican":           "co",
				"Croatian":           "hr",
				"Czech":              "cs",
				"Danish":             "da",
				"Dutch":              "nl",
				"English":            "en",
				"Esperanto":          "eo",
				"Estonian":           "et",
				"Filipino":           "tl",
				"Finnish":            "fi",
				"French":             "fr",
				"Frisian":            "fy",
				"Galician":           "gl",
				"Georgian":           "ka",
				"German":             "de",
				"Greek":              "el",
				"Gujarati":           "gu",
				"Haitian Creole":     "ht",
				"Hausa":              "ha",
				"Hawaiian":           "haw",
				"Hebrew":             "iw",
				"Hindi":              "hi",
				"Hmong":              "hmn",
				"Hungarian":          "hu",
				"Icelandic":          "is",
				"Igbo":               "ig",
				"Indonesian":         "class",
				"Irish":              "ga",
				"Italian":            "it",
				"Japanese":           "ja",
				"Javanese":           "jw",
				"Kannada":            "kn",
				"Kazakh":             "kk",
				"Khmer":              "km",
				"Kinyarwanda":        "rw",
				"Korean":             "ko",
				"Kurdish (Kurmanji)": "ku",
				"Kyrgyz":             "ky",
				"Lao":                "lo",
				"Latin":              "la",
				"Latvian":            "lv",
				"Lithuanian":         "lt",
				"Luxembourgish":      "lb",
				"Macedonian":         "mk",
				"Malagasy":           "mg",
				"Malay":              "ms",
				"Malayalam":          "ml",
				"Maltese":            "mt",
				"Maori":              "mi",
				"Marathi":            "mr",
				"Mongolian":          "mn",
				"Myanmar (Burmese)":  "my",
				"Nepali":             "ne",
				"Norwegian":          "no",
				"Odia (Oriya)":       "or",
				"Pashto":             "ps",
				"Persian":            "fa",
				"Polish":             "pl",
				"Portuguese":         "pt",
				"Punjabi":            "pa",
				"Romanian":           "ro",
				"Russian":            "ru",
				"Samoan":             "sm",
				"Scots Gaelic":       "gd",
				"Serbian":            "sr",
				"Sesotho":            "st",
				"Shona":              "sn",
				"Sindhi":             "sd",
				"Sinhala":            "si",
				"Slovak":             "sk",
				"Slovenian":          "sl",
				"Somali":             "so",
				"Spanish":            "es",
				"Sundanese":          "su",
				"Swahili":            "sw",
				"Swedish":            "sv",
				"Tajik":              "tg",
				"Tamil":              "ta",
				"Tatar":              "tt",
				"Telugu":             "te",
				"Thai":               "th",
				"Turkish":            "tr",
				"Turkmen":            "tk",
				"Ukrainian":          "uk",
				"Urdu":               "ur",
				"Uyghur":             "ug",
				"Uzbek":              "uz",
				"Vietnamese":         "vi",
				"Welsh":              "cy",
				"Xhosa":              "xh",
				"Yclassdish":         "yi",
				"Yoruba":             "yo",
				"Zulu":               "zu",
			}

		},
		func(text string) {
			return
		},
		func(text string, language string) string {
			if len(text) == 0 || len(language) == 0 {
				return ""
			} else if len(text) == 0 || len(language) == 0 {
				return ""
			}

			if language == "auto" {
				language = "en"
			}
			paramsMap := url.Values{}
			paramsMap.Add("tl", language)
			paramsMap.Add("q", strings.TrimSpace(text))
			paramsMap.Add("client", "tw-ob")

			params := paramsMap.Encode()

			return "https://translate.google.com/translate_tts?" + params
		},
		func(text string, to_language string, from_language string) string {

			if len(from_language) == 0 {
				from_language = "auto"
			}

			paramsMap := url.Values{}
			paramsMap.Add("q", strings.TrimSpace(text))
			paramsMap.Add("sl", from_language)
			paramsMap.Add("tl", to_language)

			params := paramsMap.Encode()

			r, _ := http.Get("https://translate.google.com/m?" + params)

			doc, _ := goquery.NewDocumentFromReader(r.Body)

			return doc.Find(".result-container").Text()

		},
	}
}

var GoogleTranslateEngine = newgoogleTranslateEngineStruct()
