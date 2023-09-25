package main

func main() {
	// var card string = "Eight of Hearts"
	// go compiler infers the type using syntax below
	// card := newCard()

	// declare a slice of strings
	cards := newDeck()
	// add a new card onto the slice with append
	// cards = append(cards, newCard())

	// for loop. the range keyword provides every element in the slice
	// for _, card := range cards {
	// 	fmt.Println(card)
	// }
	cards.print()
}

// func newCard() string {
// 	return "Eight of Hearts"
// }
