package main

import (
	"fmt"
)
const englishHello = "Hello, "
const spanishHello = "Hola, "
const frenchHello = "Bonjour, "
const spanish = "Spanish"
const french = "French"

func main() {
	fmt.Println(Hello("Dear", ""))
}

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	return getGreeting(language)+name
}

func getGreeting(language string) (greetingString string)  {
	switch language {
	case french:
		greetingString = frenchHello
		break
	case spanish:
		greetingString = spanishHello
		break
	default:
		greetingString = englishHello
	}
	return
}
