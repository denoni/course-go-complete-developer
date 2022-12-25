package main

import "fmt"

type languageBot interface {
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

func (englishBot) getGreeting() string {
	return "Hi, there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b languageBot) {
	fmt.Println(b.getGreeting())
}
