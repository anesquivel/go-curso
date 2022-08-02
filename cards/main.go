package main

func main() {

	cards := newDeck()

	cards.saveToFile("my_cards")

	nwCards := newDeckFromAFile("my_cards")

	nwCards.shuffle()
	nwCards.print()

	/*fmt.Println(cards.toString())
	cards.print()

	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()*/
}
