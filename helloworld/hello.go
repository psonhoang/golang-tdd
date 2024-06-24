package main

import "fmt"

const (
 spanish = "Spanish"
 french = "French"
 vietnamese = "Vietnamese"
 englishHelloPrefix = "Hello, "
 spanishHelloPrefix =  "Hola, "
 frenchHelloPrefix = "Bonjour, "
 vietnameseHelloPrefix = "Chao, "
)


func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {	// named return value that will be instantiated in the function with a zero-value
switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case vietnamese:
		prefix = vietnameseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
