package main

import "fmt"

// create a new type, deck
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}
	suits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// print the contents of the deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

// deal out 'handSize' number of cards and return both sets
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
