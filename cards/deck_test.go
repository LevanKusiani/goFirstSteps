package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 16 {
		t.Errorf("expected deck length of 16, but got %d", len(deck))
	}

	if !contains(deck, "Ace of Spades") {
		t.Error("the deck doesn't contain 'Ace of Spades'")
	}
}

func TestSaveToDiskAndReadFromDisk(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToDisk("_decktesting")

	loadedDeck := readFromDisk("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("expected deck length of 16, but got %d", len(deck))
	}

	os.Remove("_decktesting")
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
