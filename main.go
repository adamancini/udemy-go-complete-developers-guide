package main

func main() {
	// cards := newDeckFromFile("cards.txt")
	// cards.print()

	// cards = newDeckFromFile("unreal")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
