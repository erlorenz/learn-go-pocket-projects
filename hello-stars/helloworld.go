package main

import (
	"flag"
	"fmt"
)

func main() {

	var lang string
	flag.StringVar(&lang, "lang", "en", "The code of the language to be used, e.g. en")
	flag.Parse()

	greeting := greet(language(lang))
	fmt.Println(greeting)
}

// language represents the language's code
type language string

// Phrasebook holds greeting for each language
var phrasebook = map[language]string{
	"en": "Hello world",
	"fr": "Bonjour le monde",
}

func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}

	return greeting
}
