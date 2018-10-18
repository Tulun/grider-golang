package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// Can omit eb because we don't use the value.
func (englishBot) getGreeting() string {
	// Very custom logic for generating english greeting
	return "Hello there!"
}

// Can omit sb because we don't use the value.
func (spanishBot) getGreeting() string {
	// Custom logic for generating spanish greeting
	return "Hola"
}
