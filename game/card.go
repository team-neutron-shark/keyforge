package keyforge

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
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

// Stun - Mark a creature card as stunned.
func (c *Card) Stun() {
	if strings.ToLower(c.CardType) != "creature" {
		return
	}

	c.IsStunned = true
}

// Ready a given card. If a creature card is stunned then a second call
// to this function is required in order to remove exhaustion.
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
		if card.CardNumber == cardNumber && card.Expansion == setNumber {
			return card, nil
		}
	}
	errorMessage := fmt.Sprintf("no card found with set #%d and card #%d", setNumber, cardNumber)
	return Card{}, errors.New(errorMessage)
}

// FindCardsByNumber - Find cards in a pile given a set and card number.
// This ends up being useful mostly for detecting mavericks in a pile since
// mavericks are assigned unique card IDs in the Vault.
func FindCardsByNumber(cards []Card, setNumber int, cardNumber int) ([]Card, error) {
	totalCards := []Card{}

	for _, card := range cards {
		if card.CardNumber == cardNumber && card.Expansion == setNumber {
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

// FindCardsByHouse - Return an array of cards filtered by house name.
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

// GetHouses - Return an array of houses within a card pile.
// Houses are represented as strings.
func GetHouses(cards []Card) []string {
	houses := []string{}

	for _, card := range cards {
		if !HouseExists(houses, card.House) {
			houses = append(houses, card.House)
		}
	}

	return houses
}

// Shuffle - Shuffle a given card pile.
func Shuffle(cards []Card) []Card {
	for i := range cards {
		j := rand.Intn(len(cards))
		cards[i], cards[j] = cards[j], cards[i]
	}

	return cards
}

// RemoveCard - Removes a card from a given card pile.
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

// PopCard - Treats a card pile as if it were a stack and pops the top
// card of the stack off of the pile. This function is primarily useful
// for simulating draw operations, such as drawing cards from the draw
// pile into a player's hand.
func PopCard(cards []Card) []Card {
	i := len(cards) - 1
	return append(cards[:i], cards[i+1:]...)
}

// AddCard - Treats a card pile as if it were a stack and pushes a card
// onto the top of a pile. This function is primarily used in conjunction
// with PopCard() or RemoveCard() to transfer cards from one card pile to
// another card pile.
func AddCard(cards []Card, addCard Card) []Card {
	return append(cards, addCard)
}

// ChooseRandomCard - Chooses a random card from a card pile and returns
// the card object. This function is useful primarily for cards which have
// a "use" ability which requires a player to discard a random card from
// their hand.
func ChooseRandomCard(cards []Card) (int, Card) {
	i := rand.Intn(len(cards))
	return i, cards[i]
}

// DrawCard - Simulate drawing a card from one card pile into another
// card pile. This function is primarily used to simulate draws from the
// draw pile into a player's hand.
func DrawCard(source []Card, destination []Card) ([]Card, []Card) {
	// Draw the top card
	card := source[len(source)-1]
	source = RemoveCard(source, card)
	destination = AddCard(destination, card)
	return source, destination
}

// CompareCardOrder - This function inspects two Card arrays to determine
// whether or not the card orders match. Returns false in the event card
// orders do not match and true when they do.
func CompareCardOrder(original []Card, comparison []Card) bool {
	for i := range original {
		if original[i] != comparison[i] {
			return false
		}
	}
	return true
}

// CardExists - Check for a card within a card pile. This function works
// by checking card IDs within the card pile and will detect core set cards.
// In order to detect mavericks cards must be detected by set and card number.
func CardExists(cards []Card, card Card) bool {
	for _, indexCard := range cards {
		if card.ID == indexCard.ID {
			return true
		}
	}

	return false
}

func PrependCard(cards []Card, card Card) []Card {
	newCards := []Card{}

	newCards = append(newCards, card)
	newCards = append(newCards, cards...)

	return newCards
}

func SortCardsByNumber(cards []Card) []Card {
	size := len(cards)

	if size < 1 {
		return cards
	}

	for !BubbleSortByExpansionNumber(cards) {
		BubbleSortByExpansionNumber(cards)
	}

	for !BubbleSortByCardNumber(cards) {
		BubbleSortByCardNumber(cards)
	}

	return cards
}

func BubbleSortByCardNumber(cards []Card) bool {
	sorted := true

	if len(cards) < 2 {
		return sorted
	}

	for i := range cards {

		// Prevent our code from reading outside the bounds of the slice.
		if i+1 > len(cards)-1 {
			break
		}

		// If our left-most value is greater than our right-most value then
		// swap them. Mark the sorted variable to false to perform another
		// pass over the slice.
		if cards[i].CardNumber > cards[i+1].CardNumber {
			cards[i], cards[i+1] = cards[i+1], cards[i]
			sorted = false
		}
	}

	return sorted
}

func BubbleSortByExpansionNumber(cards []Card) bool {
	sorted := true

	if len(cards) < 2 {
		return sorted
	}

	for i := range cards {

		// Prevent our code from reading outside the bounds of the slice.
		if i+1 > len(cards)-1 {
			break
		}

		// If our left-most value is greater than our right-most value then
		// swap them. Mark the sorted variable to false to perform another
		// pass over the slice.
		if cards[i].Expansion > cards[i+1].Expansion {
			cards[i], cards[i+1] = cards[i+1], cards[i]
			sorted = false
		}
	}

	return sorted
}

func LoadCardsFromFile(filename string) ([]Card, error) {
	cards := []Card{}
	buffer := bytes.Buffer{}

	file, e := os.Open(filename)

	if e != nil {
		return cards, e
	}

	buffer.ReadFrom(file)

	if e != nil {
		return cards, e
	}

	e = json.Unmarshal(buffer.Bytes(), &cards)

	if e != nil {
		return cards, e
	}

	return cards, nil
}
