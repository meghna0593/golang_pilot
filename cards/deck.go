package main

import (
	"fmt"
	"math/rand"
	"time"

	// "io/ioutil"
	"os"
	"strings"
)

// Deck type in Go behaves similar to a Deck class in Python

// Create a new type of 'deck'
// which is a slice of strings
type deck []string //extends all the behaviour of a slice of string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// loop through the deck of cards and print them
func (d deck) print() { //receiver - by convention a receiver will have 1 or 2 abbr. chars; here d is a receiver of type deck
	for i, card := range d { //in python d would've been self
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string { //helper receiver method
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error { // return type is 'error' which is returned by WriteFile method

	// deprecated
	// return ioutil.WriteFile(filename, []byte(d.toString()), 0666) // anyone can write or read this file

	return os.WriteFile(filename, []byte(d.toString()), 0666) // anyone can write or read this file

}

func newDeckFromFile(filename string) deck {
	// deprecated
	// bs, err := ioutil.ReadFile(filename) // byteslice and error returned

	bs, err := os.ReadFile(filename) // byteslice and error returned
	if err != nil {
		// log the error and entirely quit the program

		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Convert []byte to string to []string to deck
	slice_of_strings := strings.Split(string(bs), ",")
	return deck(slice_of_strings)

}

func (d deck) shuffle() { // using rand.Intn
	// rand.Seed(0) // deprecated; Providing a seed of the same value will end up with the same random sequence
	for idx := range d {
		newPosition := rand.Intn(len(d) - 1) //psuedo random generator
		d[idx], d[newPosition] = d[newPosition], d[idx]

	}
}

func (d deck) shuffle_v2() { // using rand.NewSource
	source := rand.NewSource(time.Now().UnixNano()) // using current time as the seed value
	r := rand.New(source)

	for idx := range d {
		newPosition := r.Intn(len(d) - 1)
		d[idx], d[newPosition] = d[newPosition], d[idx]

	}
}
