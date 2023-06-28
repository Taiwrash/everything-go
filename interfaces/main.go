package main

import "fmt"

type bot interface {
	setGreetings() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreetings(eb)
	printGreetings(sb)

}

func printGreetings(b bot) {
	fmt.Println(b.setGreetings())
}

func (englishBot) setGreetings() string {
	return "Hello there!"
}

func (spanishBot) setGreetings() string {
	return "Holla!"
}
