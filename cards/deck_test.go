package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 20 {
		t.Errorf("Expected deck to contain 20 cards, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spaces, but got %v", d[0])
	}

	if d[len(d)-1] != "Five of Clubs" {
		t.Errorf("Expected first card to be Five of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	testFile := "_decktesting"
	os.Remove(testFile)

	deck := newDeck()
	deck.saveToFile(testFile)

	loadedDeck := newDeckFromFile(testFile)

	if len(loadedDeck) != 20 {
		t.Errorf("Expected deck to contain 20 cards, but got %v", len(deck))
	}

	os.Remove(testFile)
}
