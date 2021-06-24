package integers

import "fmt"

const farsi = "Farsi"
const french = "French"

const englishHelloPrefix = "Hello, "
const farsiHelloPrefix = "Salam, "
const frenchHelloPrefix = "Bonjour, "

const exclamationPoint = "!"
const emptyName = "World"

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case farsi:
		prefix = farsiHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

// Hello returns a personalised greeting in a given language
func Hello(name, language string) string {

	if name == "" {
		name = emptyName
	}

	return greetingPrefix(language) + name + exclamationPoint
}

func main() {
	fmt.Println(Hello("", ""))
}
