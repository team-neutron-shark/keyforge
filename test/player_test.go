package tests

import (
	keyforge "keyforge/game"
	"testing"
)

func TestPlayerSetDeck(t *testing.T) {
	player := keyforge.NewPlayer()

	deck, e := keyforge.LoadDeckFromFile("test_data/test_deck.json")

	if e != nil {
		t.Error(e.Error())
	}

	player.SetDeck(deck)

	if !keyforge.CompareCardOrder(deck.Cards, player.PlayerDeck.Cards) {
		t.Error("Player deck does not match deck on file.")
	}
}

func TestPlayerDrawCard(t *testing.T) {
	player := keyforge.NewPlayer()

	deck, e := keyforge.LoadDeckFromFile("test_data/test_deck.json")

	if e != nil {
		t.Error(e.Error())
	}

	player.SetDeck(deck)
	player.DrawCard()

	if len(player.HandPile) != 1 {
		t.Errorf("Hand contains %d cards! Should contain 1.", len(player.HandPile))
	}
}

func TestPlayerDrawHand(t *testing.T) {
	player := keyforge.NewPlayer()

	deck, e := keyforge.LoadDeckFromFile("test_data/test_deck.json")

	if e != nil {
		t.Error(e.Error())
	}

	player.SetDeck(deck)
	player.DrawHand()

	if len(player.HandPile) != 6 {
		t.Errorf("Hand contains %d cards! Should contain 6.", len(player.HandPile))
	}

}

func TestPlayerShuffleDrawPile(t *testing.T) {
	player := keyforge.NewPlayer()
	copyHand := []keyforge.Card{}

	deck, e := keyforge.LoadDeckFromFile("test_data/test_deck.json")

	if e != nil {
		t.Error(e.Error())
	}

	player.SetDeck(deck)
	player.DrawHand()

	copyHand = append(copyHand, player.HandPile...)

	player.ShuffleDrawPile()

	if keyforge.CompareCardOrder(deck.Cards, copyHand) {
		t.Error("Draw pile did not shuffle correctly!")
	}
}

func TestPlayerShuffleDiscardPile(t *testing.T) {
	player := keyforge.NewPlayer()

	deck, e := keyforge.LoadDeckFromFile("test_data/test_deck.json")

	if e != nil {
		t.Error(e.Error())
	}

	player.SetDeck(deck)

	for i := 0; i < 36; i++ {
		player.DrawCard()
	}

	if len(player.HandPile) != 36 {
		t.Errorf("Hand contains %d cards! Should contain 36.", len(player.HandPile))
	}

	for i := 0; i < 36; i++ {
		player.Discard(player.PlayerDeck.Cards[0])
	}

	if len(player.HandPile) != 0 {
		t.Errorf("Hand contains %d cards! Should contain 0.", len(player.HandPile))
	}

	if len(player.DiscardPile) != 36 {
		t.Errorf("Discard pile contains %d cards! Should contain 36.", len(player.HandPile))
	}

	player.ShuffleDiscardPile()

	if len(player.DiscardPile) != 0 {
		t.Errorf("Hand contains %d cards! Should contain 0.", len(player.HandPile))
	}

	if len(player.DrawPile) != 36 {
		t.Errorf("Discard pile contains %d cards! Should contain 36.", len(player.HandPile))
	}

}
