package main

import "fmt"

const (
	SPANISH = "Spanish"
	FRENCH  = "French"

	ENGLISH_HELLO_PREFIX = "Hello, "
	SPANISH_HELLO_PREFIX = "Hola, "
	FRENCH_HELLO_PREFIX  = "Bonjour, "
)

func greetingPrefix(language string) (prefix string) {
	switch language {
	case FRENCH:
		prefix = FRENCH_HELLO_PREFIX
	case SPANISH:
		prefix = SPANISH_HELLO_PREFIX
	default:
		prefix = ENGLISH_HELLO_PREFIX
	}
	return
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func main() {
	fmt.Println("Hello world")
}
