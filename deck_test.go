package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	expectedLength := 52
	fistCard := d[0]
	lastCard := d[len(d) - 1]

	if len(d) != expectedLength {
		t.Errorf("Expected deck length of %v, but got %v", expectedLength, len(d))
	}

	if fistCard != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", fistCard)
	}

	if lastCard != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got %v", lastCard)
	}
}

