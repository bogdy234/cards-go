package main

func main() {
	// cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))
	// fmt.Println(cards.toString())
	// cards.saveToFile("cards.txt")
	newDeckFromFile("cards.txt").print()
}
