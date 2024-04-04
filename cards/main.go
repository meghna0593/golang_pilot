package main

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades" // Go compiler will figure out that card will contain a string
	// card = "Five of Diamonds"

	// Files in the same package need not be imported

	// card := newCard()
	// fmt.Println(card)

	// without type
	// cards := []string{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// with type
	// cards := deck{"Ace of Diamonds", newCard()} // cards is of type deck
	// cards = append(cards, "Six of Spades")

	cards := newDeck()
	// cards.print()

	// Deal method that returns multiple values
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	// Type conversion to write into file
	// fmt.Println(cards.toString())

	// Saved the deck to file
	// cards.saveToFile("myCards.txt")

	// Read deck from file
	// cards := newDeckFromFile("myCards.txt")
	// cards.print()

	// Shuffle deck
	// cards.shuffle()
	cards.shuffle_v2()
	cards.print()

}

// func newCard() string {
// 	return "Five of Diamonds"
// }
