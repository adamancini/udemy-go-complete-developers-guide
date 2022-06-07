package main

import "fmt"

func main() {
	cards := newDeck()

	cards.saveToFile("cards.txt")
	hand, _ := deal(cards, 5)

	fmt.Print(hand.toString())
}
