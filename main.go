package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	// it returns a new slice that we pass to the variable cards
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
