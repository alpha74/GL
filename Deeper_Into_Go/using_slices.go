package main

import "fmt"

func main() {
	// Slice declaration
	cards := []string{newCard(), "Ace of Spades"}

	// Adding element. append() does not modify existing slice
	cards = append(cards, "Six of Club")

	// Print cards
	fmt.Println(cards)

	// Iterate on slice
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
