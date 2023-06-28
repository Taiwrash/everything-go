package main

func main() {
	cards := read("cards.txt")
	picked, remaining := deal(cards, 5)
	picked.print()
	remaining.print()
	cards.shuffle()
	cards.save("hello.in")
	cards.print()
}
