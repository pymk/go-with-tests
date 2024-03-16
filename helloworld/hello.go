package main

import "fmt"

const english = "English"
const french = "French"
const japanese = "Japanese"

const englishHelloPrefix = "Hello"
const frenchHelloPrefix = "Bonjour"
const japaneseHelloPrefix = "こんにちは"

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return fmt.Sprintf("%s, %s!", greetingPrefix(language), name)
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case japanese:
		prefix = japaneseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("", ""))
}
