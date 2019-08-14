package main

import "fmt"

const french = "french"
const spanish = "spanish"

const helloPrefixEnglish = "Hello, "
const helloPrefixFrench = "Bonjour, "
const helloPrefixSpanish = "Hola, "

func GreetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = helloPrefixFrench
	case spanish:
		prefix = helloPrefixSpanish
	default:
		prefix = helloPrefixEnglish
	}
	return
}

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return GreetingPrefix(language) + name + "!"
}

func main() {
	fmt.Println(Hello("David", "french"))
}
