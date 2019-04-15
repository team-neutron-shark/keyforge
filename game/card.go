package keyforge

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

// Card - this type represents a card, both cards used within the game as well
// as the Vault API.
type Card struct {
	ID          string `json:"id"`
	CardTitle   string `json:"card_title"`
	House       string `json:"house"`
	CardType    string `json:"card_type"`
	FrontImage  string `json:"front_image"`
	CardText    string `json:"card_text"`
	Traits      string `json:"traits"`
	Amber       int    `json:"amber"`
	Power       int    `json:"power"`
	Armor       int    `json:"armor"`
	Rarity      string `json:"rarity"`
	FlavorText  string `json:"flavor_text"`
	CardNumber  int    `json:"card_number"`
	Expansion   int    `json:"expansion"`
	IsMaverick  bool   `json:"is_maverick"`
	IsExhausted bool   `json:"-"`
	IsStunned   bool   `json:"-"`
	PowerBonus  int    `json:"-"`
	ArmorBonus  int    `json:"-"`
}

func (c *Card) Stun() {
	if strings.ToLower(c.CardType) != "creature" {
		return
	}

	c.IsStunned = true
}

func (c *Card) Ready() {
	// If stunned, remove stun, but do not ready.
	if c.IsStunned {
		c.IsStunned = false
		return
	}

	if c.IsExhausted {
		c.IsExhausted = false
	}
}

// PrettyPrint - Used to debug card data without making your eyes bleed.
func (c *Card) PrettyPrint() {
	fmt.Println("Card Title: ", c.CardTitle)
	fmt.Println("Card House: ", c.House)
	fmt.Println("Card Amber: ", c.Amber)
}

// GetActionCards - Return all action cards from a given pile.
func GetActionCards(cards []Card) []Card {
	actions := []Card{}

	for _, card := range cards {
		if strings.ToLower(card.CardType) == "action" {
			actions = append(actions, card)
		}
	}

	return actions
}

// GetArtifactCards - Return all artifact cards from a given pile.
func GetArtifactCards(cards []Card) []Card {
	artifacts := []Card{}

	for _, card := range cards {
		if strings.ToLower(card.CardType) == "artifact" {
			artifacts = append(artifacts, card)
		}
	}

	return artifacts
}

// GetCreatureCards - Return all creature cards from a given pile.
func GetCreatureCards(cards []Card) []Card {
	creatures := []Card{}

	for _, card := range cards {
		if strings.ToLower(card.CardType) == "creature" {
			creatures = append(creatures, card)
		}
	}

	return creatures
}

// GetUpgradeCards - Return all upgrade cards from a given pile.
func GetUpgradeCards(cards []Card) []Card {
	upgrades := []Card{}

	for _, card := range cards {
		if strings.ToLower(card.CardType) == "upgrade" {
			upgrades = append(upgrades, card)
		}
	}

	return upgrades
}

// FindCardByID - Find a card in a pile given a card ID.
func FindCardByID(cards []Card, cardID string) (Card, error) {
	for _, card := range cards {
		if card.ID == cardID {
			return card, nil
		}
	}
	errorMessage := fmt.Sprintf("no card found with ID %s", cardID)
	return Card{}, errors.New(errorMessage)
}

// FindCardsByID - Find a cards in a pile given a card ID.
func FindCardsByID(cards []Card, cardID string) ([]Card, error) {
	totalCards := []Card{}

	for _, card := range cards {
		if card.ID == cardID {
			totalCards = append(totalCards, card)
		}
	}

	if len(totalCards) == 0 {
		errorMessage := fmt.Sprintf("no cards found with ID %s", cardID)
		return totalCards, errors.New(errorMessage)
	}

	return totalCards, nil
}

// FindCardByNumber - Find a card in a pile given a set and card number.
// This ends up being useful mostly for detecting mavericks in a pile.
func FindCardByNumber(cards []Card, setNumber int, cardNumber int) (Card, error) {
	for _, card := range cards {
		if card.CardNumber == cardNumber && card.CardNumber == setNumber {
			return card, nil
		}
	}
	errorMessage := fmt.Sprintf("no card found with set #%d and card #%d", setNumber, cardNumber)
	return Card{}, errors.New(errorMessage)
}

// FindCardByNumber - Find a cards in a pile given a set and card number.
// This ends up being useful mostly for detecting mavericks in a pile.
func FindCardsByNumber(cards []Card, setNumber int, cardNumber int) ([]Card, error) {
	totalCards := []Card{}

	for _, card := range cards {
		if card.CardNumber == cardNumber && card.CardNumber == setNumber {
			totalCards = append(totalCards, card)
		}
	}

	if len(totalCards) == 0 {
		errorMessage := fmt.Sprintf("no card found with set #%d and card #%d", setNumber, cardNumber)
		return totalCards, errors.New(errorMessage)
	}

	return totalCards, nil
}

// GetTotalAmber - Count the amount of amber in a given card pile. This
// function only calculates the total number of amber icons in a card
// pile; it will not calculate gained/stolen amber based on board state.
func GetTotalAmber(cards []Card) int {
	amber := 0

	for _, card := range cards {
		amber += card.Amber
	}

	return amber
}

// GetTotalArmor - Count the amount of armor present in a card pile.
func GetTotalArmor(cards []Card) int {
	armor := 0

	for _, card := range cards {
		armor += card.Armor
	}

	return armor
}

// GetTotalCreaturePower - Count the total creature power in a card pile.
func GetTotalCreaturePower(cards []Card) int {
	power := 0

	for _, card := range cards {
		power += card.Power
	}

	return power
}

// GetMaximumCreaturePower - Returns the highest power value for a creature
// in a given card pile.
func GetMaximumCreaturePower(cards []Card) int {
	power := 0

	for _, card := range cards {
		if card.Power > power {
			power = card.Power
		}
	}

	return power
}

// GetMinimumCreaturePower - Returns the lowest power value for a creature
// in a given card pile.
func GetMinimumCreaturePower(cards []Card) int {
	power := 0

	if len(cards) > 0 {
		power = cards[0].Power

		for _, card := range cards {
			if card.Power < power {
				power = card.Power
			}
		}
	}

	return power
}

func FindCardsByHouse(cards []Card, house string) ([]Card, error) {
	totalCards := []Card{}

	for _, card := range cards {
		if strings.ToLower(house) == strings.ToLower(card.House) {
			totalCards = append(totalCards, card)
		}
	}

	if len(totalCards) == 0 {
		errorMessage := fmt.Sprintf("no card found with house %s", house)
		return totalCards, errors.New(errorMessage)
	}

	return totalCards, nil
}

func GetHouses(cards []Card) []string {
	houses := []string{}

	for _, card := range cards {
		if !HouseExists(&houses, card.House) {
			houses = append(houses, card.House)
		}
	}

	return houses
}

func Shuffle(cards []Card) []Card {
	for i := range cards {
		j := rand.Intn(len(cards))
		cards[i], cards[j] = cards[j], cards[i]
	}

	return cards
}

func RemoveCard(cards []Card, removeCard Card) []Card {
	returnCards := []Card{}
	found := false

	if !CardExists(cards, removeCard) {
		return returnCards
	}

	for _, card := range cards {
		if card.ID != removeCard.ID || found {
			returnCards = append(returnCards, card)
		}
		if card.ID == removeCard.ID && !found {
			found = true
		}
	}

	return returnCards
}

func PopCard(cards []Card) []Card {
	i := len(cards) - 1
	return append(cards[:i], cards[i+1:]...)
}

func AddCard(cards []Card, addCard Card) []Card {
	return append(cards, addCard)
}

func ChooseRandomCard(cards []Card) (int, Card) {
	i := rand.Intn(len(cards))
	return i, cards[i]
}

func DrawCard(source []Card, destination []Card) ([]Card, []Card) {
	// Draw the top card
	card := source[len(source)-1]
	source = RemoveCard(source, card)
	destination = AddCard(destination, card)
	return source, destination
}

func CompareCardOrder(original []Card, comparison []Card) bool {
	for i := range original {
		if original[i] != comparison[i] {
			return false
		}
	}
	return true
}

func CardExists(cards []Card, card Card) bool {
	for _, indexCard := range cards {
		if card == indexCard {
			return true
		}
	}

	return false
}
