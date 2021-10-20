package main

import (
	"os"
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

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")

	deck := newDeck()
	deck.saveToFile("_deckTesting")

	loadedDeck := newDeckFromFile("_deckTesting")

	deckLength := len(loadedDeck)
	expectedLength := 52


	if deckLength != expectedLength {
		t.Errorf("Expected %v cards in deck, but got %v", expectedLength, deckLength)
	}

	os.Remove("_deckTesting")
}