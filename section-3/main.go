package main

import (
	"fmt"
)

func main() {
	// These are equivalent declarations
	// var card string = "Ace of Spades"
	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
