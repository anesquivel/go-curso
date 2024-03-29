package main

import (
	"os"
	"testing"
)

func TestNewD(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected a deck of 16 but got: %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but go %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected first card of Four of Clubs, but go %v", d[len(d)-1])

	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromAFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in the deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
