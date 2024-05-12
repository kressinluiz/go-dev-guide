package main

import "fmt"

func main() {
	//var card string = "Ace of Spades" o tipo da variável pode ser inferido nesse caso
	card := "Ace of Spades"   // Forma simplificada. := apenas para definir uma nova variável
	card = "Five of Diamonds" // para redefinir o valor não é necessário

	var deckSize int
	deckSize = 52

	fmt.Println(card)
	fmt.Println(deckSize)
}
