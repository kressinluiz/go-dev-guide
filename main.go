package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	// it returns a new slice that we pass to the variable cards
	cards = append(cards, "Six of Spades")

	//range is a keyword that we use when we want to iterate over every element of a slice
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
