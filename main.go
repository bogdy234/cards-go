package main

import "fmt"

func main() {
	// cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))
	// fmt.Println(cards.toString())
	// cards.saveToFile("cards.txt")
	deck := newDeckFromFile("cards.txt")
	fmt.Println(deck)
	deck.shuffle()
	fmt.Println("------")
	fmt.Println(deck)
}
