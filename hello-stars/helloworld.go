package main

import "fmt"

func main() {
	greeting := greet("en")
	fmt.Println(greeting)
}

// language representst the language's code
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
