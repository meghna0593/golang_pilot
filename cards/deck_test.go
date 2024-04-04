package main

import "testing"

func TestNewDeck(t *testing.T) { //automatically called by Go test runner using t as the test handler to report test failures
	d := newDeck()

	if len(d) != 20 { //depending on the deck created in deck.go
		t.Errorf("Expected deck length of 20, but got %v", len(d))
	}
}
