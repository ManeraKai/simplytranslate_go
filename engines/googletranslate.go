package engines

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type googleTranslateEngineStruct struct{}

func (self googleTranslateEngineStruct) GetSupportedLanguages() map[string]string {
	return map[string]string{
		"af":    "Afrikaans",
		"sq":    "Albanian",
		"am":    "Amharic",
		"ar":    "Arabic",
		"hy":    "Armenian",
		"az":    "Azerbaijani",
		"eu":    "Basque",
		"be":    "Belarusian",
		"bn":    "Bengali",
		"bs":    "Bosnian",
		"bg":    "Bulgarian",
		"ca":    "Catalan",
		"ceb":   "Cebuano",
		"ny":    "Chichewa",
		"zh-CN": "Chinese",
		"co":    "Corsican",
		"hr":    "Croatian",
		"cs":    "Czech",
		"da":    "Danish",
		"nl":    "Dutch",
		"en":    "English",
		"eo":    "Esperanto",
		"et":    "Estonian",
		"tl":    "Filipino",
		"fi":    "Finnish",
		"fr":    "French",
		"fy":    "Frisian",
		"gl":    "Galician",
		"ka":    "Georgian",
		"de":    "German",
		"el":    "Greek",
		"gu":    "Gujarati",
		"ht":    "Haitian Creole",
		"ha":    "Hausa",
		"haw":   "Hawaiian",
		"iw":    "Hebrew",
		"hi":    "Hindi",
		"hmn":   "Hmong",
		"hu":    "Hungarian",
		"is":    "Icelandic",
		"ig":    "Igbo",
		"class": "Indonesian",
		"ga":    "Irish",
		"it":    "Italian",
		"ja":    "Japanese",
		"jw":    "Javanese",
		"kn":    "Kannada",
		"kk":    "Kazakh",
		"km":    "Khmer",
		"rw":    "Kinyarwanda",
		"ko":    "Korean",
		"ku":    "Kurdish (Kurmanji)",
		"ky":    "Kyrgyz",
		"lo":    "Lao",
		"la":    "Latin",
		"lv":    "Latvian",
		"lt":    "Lithuanian",
		"lb":    "Luxembourgish",
		"mk":    "Macedonian",
		"mg":    "Malagasy",
		"ms":    "Malay",
		"ml":    "Malayalam",
		"mt":    "Maltese",
		"mi":    "Maori",
		"mr":    "Marathi",
		"mn":    "Mongolian",
		"my":    "Myanmar (Burmese)",
		"ne":    "Nepali",
		"no":    "Norwegian",
		"or":    "Odia (Oriya)",
		"ps":    "Pashto",
		"fa":    "Persian",
		"pl":    "Polish",
		"pt":    "Portuguese",
		"pa":    "Punjabi",
		"ro":    "Romanian",
		"ru":    "Russian",
		"sm":    "Samoan",
		"gd":    "Scots Gaelic",
		"sr":    "Serbian",
		"st":    "Sesotho",
		"sn":    "Shona",
		"sd":    "Sindhi",
		"si":    "Sinhala",
		"sk":    "Slovak",
		"sl":    "Slovenian",
		"so":    "Somali",
		"es":    "Spanish",
		"su":    "Sundanese",
		"sw":    "Swahili",
		"sv":    "Swedish",
		"tg":    "Tajik",
		"ta":    "Tamil",
		"tt":    "Tatar",
		"te":    "Telugu",
		"th":    "Thai",
		"tr":    "Turkish",
		"tk":    "Turkmen",
		"uk":    "Ukrainian",
		"ur":    "Urdu",
		"ug":    "Uyghur",
		"uz":    "Uzbek",
		"vi":    "Vietnamese",
		"cy":    "Welsh",
		"xh":    "Xhosa",
		"yi":    "Yclassdish",
		"yo":    "Yoruba",
		"zu":    "Zulu",
	}

}

func (self googleTranslateEngineStruct) Translate(text, from, to string) string {

	if len(from) == 0 {
		from = "auto"
	}

	paramsMap := url.Values{}
	paramsMap.Add("q", strings.TrimSpace(text))
	paramsMap.Add("sl", from)
	paramsMap.Add("tl", to)

	params := paramsMap.Encode()

	r, httpError := http.Get("https://translate.google.com/m?" + params)

	if httpError != nil {
		fmt.Println("httpError:", httpError)
	}

	doc, goqueryError := goquery.NewDocumentFromReader(r.Body)

	if goqueryError != nil {
		fmt.Println("goqueryError", goqueryError)
	}

	return doc.Find(".result-container").Text()

}

func (self googleTranslateEngineStruct) Name() string {
	return "google"
}

var GoogleTranslateEngine = func() *googleTranslateEngineStruct {
	return &googleTranslateEngineStruct{}
}()
