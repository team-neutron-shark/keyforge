package tests

import (
	keyforge "keyforge/game"
	"math/rand"
	"testing"
)

func TestCardGetActionCards(t *testing.T) {
	deck, e := keyforge.LoadDeckFromFile("test_data/test_deck.json")

	if e != nil {
		t.Error(e.Error())
	}

	actions := keyforge.GetActionCards(deck.Cards)

	if len(actions) != 16 {
		t.Errorf("Deck contains %d actions cards! Should be 16 cards.", len(actions))
	}
}

func TestCardGetArtifactCards(t *testing.T) {
	deck, e := keyforge.LoadDeckFromFile("test_data/test_deck.json")

	if e != nil {
		t.Error(e.Error())
	}

	artifacts := keyforge.GetArtifactCards(deck.Cards)

	if len(artifacts) != 3 {
		t.Errorf("Deck contains %d artifact cards! Should be 3 cards.", len(artifacts))
	}
}

func TestCardGetCreatureCards(t *testing.T) {
	deck, e := keyforge.LoadDeckFromFile("test_data/test_deck.json")

	if e != nil {
		t.Error(e.Error())
	}

	creatures := keyforge.GetCreatureCards(deck.Cards)

	if len(creatures) != 16 {
		t.Errorf("Deck contains %d artifact cards! Should be 16 cards.", len(creatures))
	}
}

func TestCardFindCardByID(t *testing.T) {
	deck, e := keyforge.LoadDeckFromFile("test_data/test_deck.json")

	if e != nil {
		t.Error(e.Error())
	}

	// Deck contains two Coward's End cards, should be able to find one
	_, e = keyforge.FindCardByID(deck.Cards, "d438faa9-7920-437a-8d1c-682fade5d350")

	if e != nil {
		t.Error(e.Error())
	}
}

func TestCardFindCardsByID(t *testing.T) {
	deck, e := keyforge.LoadDeckFromFile("test_data/test_deck.json")

	if e != nil {
		t.Error(e.Error())
	}

	// Deck contains two Coward's End cards
	cards, e := keyforge.FindCardsByID(deck.Cards, "d438faa9-7920-437a-8d1c-682fade5d350")

	if e != nil {
		t.Error(e.Error())
	}

	if len(cards) != 2 {
		t.Errorf("Deck contains %d artifact cards! Should be 2 cards.", len(cards))
	}
}

func TestCardGetUpgradeCards(t *testing.T) {
	deck, e := keyforge.LoadDeckFromFile("test_data/test_deck.json")

	if e != nil {
		t.Error(e.Error())
	}

	upgrades := keyforge.GetUpgradeCards(deck.Cards)

	if len(upgrades) != 1 {
		t.Errorf("Deck contains %d artifact cards! Should be 1 card.", len(upgrades))
	}
}

func TestCardFindCardByNumber(t *testing.T) {
	deck, e := keyforge.LoadDeckFromFile("test_data/test_deck.json")

	if e != nil {
		t.Error(e.Error())
	}

	_, e = keyforge.FindCardByNumber(deck.Cards, 341, 14)

	if e != nil {
		t.Error(e.Error())
	}
}

func TestCardFindCardsByNumber(t *testing.T) {
	deck, e := keyforge.LoadDeckFromFile("test_data/test_deck.json")

	if e != nil {
		t.Error(e.Error())
	}

	cards, e := keyforge.FindCardsByNumber(deck.Cards, 341, 14)

	if e != nil {
		t.Error(e.Error())
	}

	if len(cards) < 1 {
		t.Errorf("Could not find test card; should have 1 in the deck.")
	}

	cards, e = keyforge.FindCardsByNumber(deck.Cards, 341, 7)

	if e != nil {
		t.Error(e.Error())
	}

	if len(cards) < 2 {
		t.Errorf("Could not find test card; should have 2 in the deck.")
	}
}

func TestCardFindCardsByHouse(t *testing.T) {
	deck, e := keyforge.LoadDeckFromFile("test_data/test_deck.json")

	if e != nil {
		t.Error(e.Error())
	}

	cards, e := keyforge.FindCardsByHouse(deck.Cards, "Dis")

	if e != nil {
		t.Error(e.Error())
	}

	if len(cards) != 12 {
		t.Errorf("Deck contains %d Dis cards! Should contain 12.", len(cards))
	}
}

func TestCardGetHouses(t *testing.T) {
	deck, e := keyforge.LoadDeckFromFile("test_data/test_deck.json")

	if e != nil {
		t.Error(e.Error())
	}

	houses := keyforge.GetHouses(deck.Cards)

	if len(houses) != 3 {
		t.Errorf("There were %d houses detected in this deck! Each deck should contain 3 houses.", len(houses))
	}
}

func TestCardRemoveCard(t *testing.T) {
	deck, e := keyforge.LoadDeckFromFile("test_data/test_deck.json")

	if e != nil {
		t.Error(e.Error())
	}

	// Test deck contains two Coward's End cards, we're going to remove one
	card, e := keyforge.FindCardByID(deck.Cards, "d438faa9-7920-437a-8d1c-682fade5d350")

	if e != nil {
		t.Error(e.Error())
	}

	deck.Cards = keyforge.RemoveCard(deck.Cards, card)

	if len(deck.Cards) != 35 {
		t.Errorf("Deck contains %d cards after one card removed! Should be 35 cards.", len(deck.Cards))
	}

	cards, e := keyforge.FindCardsByID(deck.Cards, "d438faa9-7920-437a-8d1c-682fade5d350")

	if e != nil {
		t.Error(e.Error())
	}

	// Test deck should containg one Coward's End at this point.
	if len(cards) != 1 {
		t.Errorf("Detecteds %d Coward's End cards. Should be 1 remaining.", len(cards))
	}
}

func TestCardAddCard(t *testing.T) {
	cardPile := []keyforge.Card{}

	deck, e := keyforge.LoadDeckFromFile("test_data/test_deck.json")

	if e != nil {
		t.Error(e.Error())
	}

	card, e := keyforge.FindCardByID(deck.Cards, "d438faa9-7920-437a-8d1c-682fade5d350")

	if e != nil {
		t.Error(e.Error())
	}

	cardPile = keyforge.AddCard(cardPile, card)

	if len(cardPile) != 1 {
		t.Errorf("Card pile contains %d cards after one card removed! Should be 1 card.", len(cardPile))
	}

	if cardPile[0].ID != "d438faa9-7920-437a-8d1c-682fade5d350" {
		t.Errorf("The wrong card was somehow added to the card pile. Card ID added: %s.", cardPile[0].ID)
	}
}

func TestCardShuffle(t *testing.T) {
	copyCards := []keyforge.Card{}
	deck, e := keyforge.LoadDeckFromFile("test_data/test_deck.json")

	copyCards = append(copyCards, deck.Cards...)

	if e != nil {
		t.Error(e.Error())
	}

	copyCards = keyforge.Shuffle(copyCards)

	if keyforge.CompareCardOrder(deck.Cards, copyCards) {
		t.Error("Cards have not been shuffled!")
	}
}

func TestCardPopCard(t *testing.T) {
	deck, e := keyforge.LoadDeckFromFile("test_data/test_deck.json")

	if e != nil {
		t.Error(e.Error())
	}

	card := deck.Cards[len(deck.Cards)-1]

	deck.Cards = keyforge.PopCard(deck.Cards)

	compareCard := deck.Cards[len(deck.Cards)-1]

	if card.ID == compareCard.ID {
		t.Error("Card was not popped off the pile!")
	}
}

func TestCardDrawCard(t *testing.T) {
	hand := []keyforge.Card{}
	deck, e := keyforge.LoadDeckFromFile("test_data/test_deck.json")

	if e != nil {
		t.Error(e.Error())
	}

	deck.Cards, hand = keyforge.DrawCard(deck.Cards, hand)

	if len(hand) != 1 {
		t.Errorf("There are %d cards in the hand! Should be 1 card.", len(hand))
	}

	if hand[0].ID != "bec84d69-68f0-456c-a7bd-9f1e94d55a22" {
		t.Errorf("Should have drawn Experimental Therapy, drew %s instead!", hand[0].CardTitle)
	}
}

func TestCardCompareCardOrder(t *testing.T) {
	copyCards := []keyforge.Card{}
	deck, e := keyforge.LoadDeckFromFile("test_data/test_deck.json")

	if e != nil {
		t.Error(e.Error())
	}

	copyCards = append(copyCards, deck.Cards...)

	if !keyforge.CompareCardOrder(deck.Cards, copyCards) {
		t.Error("Card orders do not match!")
	}

	copyCards = keyforge.Shuffle(copyCards)

	if keyforge.CompareCardOrder(deck.Cards, copyCards) {
		t.Error("Card orders match after shuffling!")
	}
}

func TestCardCardExists(t *testing.T) {
	deck, e := keyforge.LoadDeckFromFile("test_data/test_deck.json")

	if e != nil {
		t.Error(e.Error())
	}

	card, e := keyforge.FindCardByID(deck.Cards, "0ef760a3-68b9-42a9-93fa-419ea171917b")

	if e != nil {
		t.Error(e.Error())
	}

	if !keyforge.CardExists(deck.Cards, card) {
		t.Errorf("%s not found in deck!", card.CardTitle)
	}

	deck.Cards = keyforge.RemoveCard(deck.Cards, card)

	if keyforge.CardExists(deck.Cards, card) {
		t.Errorf("Card %s found after being removed from the deck!", card.CardTitle)
	}
}

func TestCardPrependCard(t *testing.T) {
	hand := []keyforge.Card{}

	deck, e := keyforge.LoadDeckFromFile("test_data/test_deck.json")

	selectedCard := deck.Cards[rand.Intn(len(deck.Cards))]

	if e != nil {
		t.Error(e.Error())
	}

	hand = keyforge.PrependCard(hand, selectedCard)

	if len(hand) != 1 {
		t.Errorf("Hand contains %d cards! Should contain 1.", len(hand))
	}

	if hand[0].ID != selectedCard.ID {
		t.Errorf("Incorrect card prepended to the array! Should be %s.", selectedCard.CardTitle)
	}

	hand = nil
	hand = keyforge.AddCard(hand, deck.Cards[0])
	hand = keyforge.AddCard(hand, deck.Cards[1])
	hand = keyforge.AddCard(hand, deck.Cards[2])

	hand = keyforge.PrependCard(hand, selectedCard)

	if len(hand) != 4 {
		t.Errorf("There are %d currently in hand! Should be 4.", len(hand))
	}

	if hand[0].ID != selectedCard.ID {
		t.Errorf("Incorrect card prepended to the array! Should be %s", selectedCard.CardTitle)
	}
}

func TestCardSortCards(t *testing.T) {
	deck, e := keyforge.LoadDeckFromFile("test_data/test_deck.json")

	if e != nil {
		t.Error(e.Error())
	}

	testSlice := keyforge.SortCardsByNumber(deck.Cards)

	for i, card := range testSlice {
		if i+1 > len(testSlice)-1 {
			break
		}

		if card.CardNumber > testSlice[i+1].CardNumber {
			t.Error("Deck was not properly sorted!")
		}
	}
}
