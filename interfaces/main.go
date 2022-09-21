package main

import "fmt"

type bot interface {
	getGreeting() string
	getMessages() []string
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	return "hello!"
}

func (spanishBot) getGreeting() string {
	return "hola!"
}

func (englishBot) getMessages() string {
	return "mes"
}
func (spanishBot) getMessages() string {
	return "mes"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}
