package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	// very custom logic for returning and english greeting
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

func main() {

}
