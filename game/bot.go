package keyforge

import (
	"fmt"
	"strings"
)

type Bot struct {
	Player
}

func NewBot() *Bot {
	bot := new(Bot)
	return bot
}

func (b *Bot) DetermineMulligan() bool {
	houses := GetHouses(b.HandPile)
	creatures := GetCreatureCards(b.HandPile)
	amber := GetTotalAmber(b.HandPile)

	fmt.Println("houses in hand:", strings.Join(houses, ", "))

	// If our hand contains 3 creatures, do not mulligan.
	if len(creatures) > 2 {
		return false
	}

	// If our hand contains 2 creatures from only 2 houses, do not mulligan.
	if len(creatures) > 1 && len(houses) == 2 {
		return false
	}

	// Got first draw, hand only contains 2 houses, do not mulligan.
	if len(b.HandPile) == 7 && len(houses) == 2 {
		return false
	}

	if amber > 2 {
		return false
	}

	return true
}

func (b *Bot) DetermineActiveHouse() string {
	houses := GetHouses(b.HandPile)
	cards := map[string][]Card{}
	maxCount := 0
	houseChoice := ""

	for _, house := range houses {
		foundCards, _ := FindCardsByHouse(b.HandPile, house)
		cards[house] = foundCards
	}

	for house, houseCards := range cards {
		if len(houseCards) > maxCount {
			maxCount = len(houseCards)
			houseChoice = house
		}
	}

	return houseChoice
}

func (b *Bot) PlayCards(house string) {
	cards, e := FindCardsByHouse(b.HandPile, house)

	if e != nil {
		fmt.Println(e)
		return
	}

	for _, card := range cards {
		b.PlayCard(card)
	}
}
