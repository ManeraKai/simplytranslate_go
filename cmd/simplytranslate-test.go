package main

import (
	"fmt"

	engines "github.com/manerakai/simplytranslate_go/engines"
)

func main() {

	fmt.Println("Google supported languages:")
	fmt.Println(engines.GetSupportedLanguages(engines.GoogleTranslateEngine))

	fmt.Println("")

	fmt.Println("Google Translating `Hello Weird World!` from `en` to `es`:")
	fmt.Println(engines.Translate("Hello Weird World!", "en", "es", engines.GoogleTranslateEngine))

	fmt.Println("")

	fmt.Println("Google Getting FullName of code `en`:")
	fmt.Println(engines.ToFullName("en", engines.GoogleTranslateEngine))

	fmt.Println("")

	fmt.Println("Google Getting LangCode of `English`")
	fmt.Println(engines.ToLangCode("English", engines.GoogleTranslateEngine))
}
