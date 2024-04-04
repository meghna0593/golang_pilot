package main

import "testing"

func TestNewDeck(t *testing.T) { //automatically called by Go test runner using t as the test handler to report test failures
	d := newDeck()

	// Test1: Make sure the length is as expected
	if len(d) != 20 { //depending on the deck created in deck.go
		t.Errorf("Expected deck length of 20, but got %v", len(d))
	}

	// Test2: Make sure the first card in the deck is Ace of Spades
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	// Test3: Make sure the last card in the deck is Five of Clubs
	if d[len(d)-1] != "Five of Clubs" {
		t.Errorf("Expected first card of Five of Clubs, but got %v", d[len(d)-1])
	}
}
