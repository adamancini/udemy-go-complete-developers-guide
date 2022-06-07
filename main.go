package main

func main() {
	cards := newDeckFromFile("cards.txt")
	cards.print()

	cards = newDeckFromFile("unreal")
	cards.print()

}
