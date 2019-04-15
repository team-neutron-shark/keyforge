package keyforge

import (
	"strings"
)

func HouseExists(array *[]string, house string) bool {
	for _, s := range *array {
		if strings.ToLower(s) == strings.ToLower(house) {
			return true
		}
	}
	return false
}

func PrepareDrawPile(player *Player) {
	player.DrawPile = nil
	player.DrawPile = append(player.DrawPile, player.PlayerDeck.Cards...)

	for i := 0; i < 10; i++ {
		Shuffle(player.DrawPile)
	}

	for i := 0; i < 6; i++ {
		player.DrawCard()
	}
}
