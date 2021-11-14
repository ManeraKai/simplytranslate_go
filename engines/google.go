package engines

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

type googleTranslateEngineStruct struct {
	Name string
}

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

	if from == "" {
		from = "auto"
	}

	paramsMap := url.Values{}

	paramsMap.Add("sl", from)
	paramsMap.Add("tl", to)
	paramsMap.Add("hl", to)
	paramsMap.Add("q", strings.TrimSpace(text))

	params := paramsMap.Encode()

	myUrl := "https://translate.googleapis.com/translate_a/single?client=gtx&ie=UTF-8&oe=UTF-8&dt=bd&dt=ex&dt=ld&dt=md&dt=rw&dt=rm&dt=ss&dt=t&dt=at&dt=qc&" + params

	println(myUrl)
	r, httpError := http.Get(myUrl)

	if httpError != nil {
		fmt.Println("httpError:", httpError)
	}

	println(r.Body)

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	bodyString := string(bodyBytes)

	firstAbstraction := gjson.Get(bodyString, "0")

	lengthy, lengthyError := strconv.Atoi(firstAbstraction.Get("#").Raw)
	if lengthyError != nil {
		log.Fatal(lengthyError)
	}

	var translation string

	for i := 0; i < lengthy; i++ {
		third := firstAbstraction.Get(strconv.Itoa(i)).Get("0").Str
		if third != "" {
			translation += third
		}
	}

	return translation

}

func (self googleTranslateEngineStruct) Tts(text, lang string) []byte {

	paramsMap := url.Values{}
	paramsMap.Add("q", strings.TrimSpace(text))
	paramsMap.Add("tl", lang)
	paramsMap.Add("client", "tw-ob")

	params := paramsMap.Encode()

	resp, httpError := http.Get("https://translate.google.com/translate_tts?" + params)

	if httpError != nil {
		fmt.Println("httpError:", httpError)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error with bytes", err)
	}

	return bodyBytes

}

var GoogleTranslateEngine = func() *googleTranslateEngineStruct {
	return &googleTranslateEngineStruct{"google"}
}()
