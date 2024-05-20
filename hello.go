package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(s string, language string) string {
	prefix := englishHelloPrefix

	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	}
	if s != "" {
		return prefix + s
	}
	return "Hello, World!"
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
