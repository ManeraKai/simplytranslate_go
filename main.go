package main

import (
	"fmt"

	s "github.com/manerakai/simplytranslate_go/simplytranslate_engines"
)

func main() {

	fmt.Println("Google supported languages:")
	fmt.Println(s.GetSupportedLanguages(s.GoogleTranslateEngine))

	fmt.Println("")

	fmt.Println("Google Translating `Hello Weird World!` from `en` to `es`:")
	fmt.Println(s.Translate("Hello Weird World!", "en", "es", s.GoogleTranslateEngine))

	fmt.Println("")

	fmt.Println("Google Getting FullName of code `en`:")
	fmt.Println(s.Utils.ToFullName("en", s.GoogleTranslateEngine))

	fmt.Println("")

	fmt.Println("Google Getting LangCode of `English`")
	fmt.Println(s.Utils.ToLangCode("English", s.GoogleTranslateEngine))
}
