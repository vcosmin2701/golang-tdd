package main

import (
	"fmt"
)

const (
	french   = "French"
	spanish  = "Spanish"
	romanian = "Romanian"

	englishHelloPrefix  = "Hello, "
	spanishHelloPrefix  = "Hola, "
	frenchHelloPrefix   = "Bonjour, "
	romanianHelloPrefix = "Salut, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func FullGreeting(name string, formal bool) string {
	if name == "" {
		name = "World"
	}

	return getFormalPrefix(formal) + name + "!"
}

// Exercise Full Greeting: FullGreeting(name, formal bool)
// Add "Good day" if formal, else "Sup, "

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case romanian:
		prefix = romanianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func getFormalPrefix(formal bool) (prefix string) {
	if formal {
		prefix = "Good day, "
	} else {
		prefix = "Sup, "
	}
	return
}

func main() {
	fmt.Println(Hello("world", "French"))
}
