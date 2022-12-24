package main

type deck []string

func main() {
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}
