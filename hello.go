package main

import (
	"fmt"
)

const (
	english            = "English"
	spanish            = "Spanish"
	french             = "French"
	englishPrefixHello = "Hello, "
	spanishPrefixHello = "Hola, "
	frenchPrefixHello  = "Bonjour, "
)

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	prefix := englishPrefixHello
	switch lang {
	case spanish:
		prefix = spanishPrefixHello
	case french:
		prefix = frenchPrefixHello
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
