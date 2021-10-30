package main

import (
	"fmt"

	"github.com/manerakai/simplytranslate_go/engines"
)

func main() {

	fmt.Println("Google supported languages:")
	fmt.Println(engines.GetSupportedLanguages("google"))

	fmt.Println("")

	fmt.Println("Google Translating `Hello Weird World!` from `en` to `es`:")
	fmt.Println(engines.Translate("Hello Weird World!", "en", "es", "google"))

	fmt.Println("")

	fmt.Println("Google Getting FullName of code `en`:")
	fmt.Println(engines.ToFullName("en", "google"))

	fmt.Println("")

	fmt.Println("Google Getting LangCode of `English`:")
	fmt.Println(engines.ToLangCode("English", "google"))

	fmt.Println("")

	fmt.Println("LibreTranslate supported languages:")
	fmt.Println(engines.GetSupportedLanguages("libre"))

	fmt.Println("")

	fmt.Println("LibreTranslate Translating `Hello Weird World!` from `en` to `es`:")
	fmt.Println(engines.Translate("Hello Weird World!", "en", "fr", "libre"))

	fmt.Println("")

	fmt.Println("LibreTranslate Detect language:")
	fmt.Println(engines.LibreTranslateEngine.DetectLanguage("Bonjour"))

	fmt.Println(engines.GetEngine("libre").GetSupportedLanguages())

}
