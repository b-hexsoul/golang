package main

import "fmt"

type deck []string

// this func returns a type of deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Hearts", "Clubs", "Diamonds", "Spades"}
	cardValues := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
