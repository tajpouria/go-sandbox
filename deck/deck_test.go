package main

import (
	"os"
	"testing"
)

const (
	deckSize         = 52
	firstCard        = "Ace Of Clubs"
	lastCard         = "King Of Spades"
	deckTestFilename = "_deckTest.txt"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	newDeckTestUtils(d, t)
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove(deckTestFilename)

	sourceDeck := newDeck()
	sourceDeck.saveToFile(deckTestFilename)

	d := newDeck()

	newDeckTestUtils(d, t)

	os.Remove(deckTestFilename)
}

func newDeckTestUtils(d deck, t *testing.T) {
	if len(d) != deckSize {
		t.Errorf("Expect deckSize of %v but get %v", deckSize, len(d))
	}

	if d[0] != firstCard {
		t.Errorf("Expect firstCard of %v but got %v", firstCard, d[0])
	}

	if d[len(d)-1] != lastCard {
		t.Errorf("Expect lastCard of %v but got %v", lastCard, d[len(d)-1])
	}
}
