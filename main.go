package main

import (
	"fmt"

	s "github.com/manerakai/simplytranslate_go/simplytranslate_engines"
)

func main() {
	fmt.Println(s.LibreTranslateEngine.Get_supported_languages())
}
