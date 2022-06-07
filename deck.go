package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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

// save the deck contents to a file on the local machine but first
// we have to convert a deck to a []string slice, then concatenate to a string,
// and finally to a []byte slice
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// should we use a receiver or pass a deck as an argument to a function?
// we'll talk about it in the future
// for argument's sake, it makes sense to make this a receiver
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// option #1 log the error and return a call to newDeck()
		// option #2 log the error and entirely quit
		fmt.Println("Error: ", err)
		// os.Exit(1) // option #2
		return newDeck() // option #1
	}

	return deck(strings.Split(string(bs), ","))
}
