package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	suitList := []string{"Spades", "Diamonds", "Clubs", "Hearts "}
	valueList := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range suitList {
		for _, value := range valueList {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for index, card := range d {
		fmt.Printf("%v -> %v \n", index, card)
	}
}

func deal(d deck, hand int) (deck, deck) {
	return d[:hand], d[hand:]
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	ss := rand.New(source)
	for i := range d {
		position := ss.Intn(len(d) - 1)
		d[i], d[position] = d[position], d[i]
	}
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func readFrom(filename string) deck {
	content, _ := ioutil.ReadFile(filename)
	strContent := strings.Split(string(content), ",")
	return deck(strContent)
}
