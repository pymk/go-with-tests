package main

import "fmt"

const (
	french   = "French"
	japanese = "Japanese"

	englishHelloPrefix  = "Hello, "
	frenchHelloPrefix   = "Bonjour, "
	japaneseHelloPrefix = "こんにちは, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case japanese:
		prefix = japaneseHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", "Japanese"))
}
