package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, ": ", card)
	}
}

func newDeck() deck {
	cards := deck{}
	suitCard := []string{"Hearts", "Clubs", "Spades", "Diamonds"}
	cardValue := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range suitCard {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, hand int) (deck, deck) {
	return d[:hand], d[hand:]
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	ss := rand.New(source)
	for index := range d {
		position := ss.Intn(len(d) - 1)
		d[index], d[position] = d[position], d[index]
	}
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) save(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func read(filename string) deck {
	content, _ := ioutil.ReadFile(filename)
	cc := strings.Split(string(content), ",")
	return deck(cc)
}
