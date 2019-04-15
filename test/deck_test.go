package tests

import (
	keyforge "keyforge/game"
	"testing"
)

var testLocation = "test_data/test_deck.json"

func TestLoadDeckFromFile(t *testing.T) {
	deck, e := keyforge.LoadDeckFromFile(testLocation)

	if e != nil {
		t.Error(e.Error())
	}

	if len(deck.Cards) != 36 {
		t.Errorf("Deck only contains %d cards!", len(deck.Cards))
	}
}
