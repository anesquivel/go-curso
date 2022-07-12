package main

func main() {
	//var card string = "Ace of Spades"
	//card := newCard()
	//card = "Five of Diamonds"
	//fmt.Println(card)

	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
