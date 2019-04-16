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

func TestPlayerDeployCreature(t *testing.T) {
	player := keyforge.NewPlayer()

	deck, e := keyforge.LoadDeckFromFile("test_data/test_deck.json")

	if e != nil {
		t.Error(e.Error())
	}

	player.SetDeck(deck)

	creatures := keyforge.GetCreatureCards(deck.Cards)

	if len(creatures) < 1 {
		t.Error("No creatures found in deck!")
	}

	player.Creatures = player.DeployCreature(creatures[0])

	if len(player.Creatures) != 1 {
		t.Errorf("There are %d creatures on the field! There should be 1 creature.", len(player.Creatures))
	}

	if creatures[0].ID != player.Creatures[0].ID {
		t.Errorf("Wrong card added to the field! Should have been %s.", creatures[0].CardTitle)
	}
}

func TestPlayerDeployCreatureLeftFlank(t *testing.T) {
	player := keyforge.NewPlayer()

	deck, e := keyforge.LoadDeckFromFile("test_data/test_deck.json")

	if e != nil {
		t.Error(e.Error())
	}

	player.SetDeck(deck)

	creatures := keyforge.GetCreatureCards(player.PlayerDeck.Cards)

	if len(creatures) < 2 {
		t.Error("Deck not suitable for this test; less than 2 unique creature cards available.")
	}

	if creatures[0].ID == creatures[len(creatures)-1].ID {
		t.Error("Cards chosen by the test were not unique!")
	}

	player.Creatures = player.DeployCreature(creatures[len(creatures)-1])
	player.Creatures = player.DeployCreatureLeftFlank(creatures[0])

	if len(player.Creatures) != 2 {
		t.Errorf("There are %d creatures on the field! Should be 2.", len(player.Creatures))
	}

	if player.Creatures[0].ID != creatures[0].ID {
		t.Errorf("Incorrect card placed at the left flank! Should be %s", creatures[0].CardTitle)
	}
}

func TestPlayerDeployCreatureRightFlank(t *testing.T) {
	player := keyforge.NewPlayer()

	deck, e := keyforge.LoadDeckFromFile("test_data/test_deck.json")

	if e != nil {
		t.Error(e.Error())
	}

	player.SetDeck(deck)

	creatures := keyforge.GetCreatureCards(player.PlayerDeck.Cards)

	if len(creatures) < 2 {
		t.Error("Deck not suitable for this test; less than 2 unique creature cards available.")
	}

	if creatures[0].ID == creatures[len(creatures)-1].ID {
		t.Error("Cards chosen by the test were not unique!")
	}

	player.Creatures = player.DeployCreature(creatures[len(creatures)-1])
	player.Creatures = player.DeployCreatureRightFlank(creatures[0])

	if len(player.Creatures) != 2 {
		t.Errorf("There are %d creatures on the field! Should be 2.", len(player.Creatures))
	}

	if player.Creatures[len(player.Creatures)-1].ID != creatures[0].ID {
		t.Errorf("Incorrect card placed at the left flank! Should be %s", creatures[0].CardTitle)
	}
}
