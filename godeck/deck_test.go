package main

import (
	"os"
	"testing"
)

const deckSize int = 52
const firstCard string = "Ace of Spades"
const lastCard string = "King of Clubs"

func TestNewDeck(t *testing.T) {
    d := newDeck()


    if(len(d) != deckSize) {
        t.Errorf("Expected deck length of %v, but got %v", deckSize, len(d))
    }

    if d[0] != firstCard {
        t.Errorf("Expected \"%v\" to be the first card in the deck, got %v instead", firstCard, d[0])
    }

    if d[len(d)-1] != lastCard {
        t.Errorf("Expected \"%v\" to be the last card in the deck, got %v instead", lastCard, d[len(d)-1])
    }
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
    os.Remove("_decktesting")

    deck := newDeck()
    deck.saveToFile("_decktesting")

    loadedDeck := newDeckFromFile("_decktesting")
    if(len(loadedDeck) != deckSize){
        t.Errorf("Expected deck length of %v, but got %v", deckSize, len(deck))
    }

    os.Remove("_decktesting")

}
