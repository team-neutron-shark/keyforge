package keyforge

import (
	"encoding/json"
	"io/ioutil"
)

type Deck struct {
	Name          string   `json:"name"`
	Expansion     int      `json:"expansion"`
	Chains        int      `json:"chains"`
	Wins          int      `json:"wins"`
	Losses        int      `json:"losses"`
	ID            string   `json:"id"`
	IsMyDeck      bool     `json:"is_my_deck"`
	Notes         []string `json:"notes"`
	IsMyFavorite  bool     `json:"is_my_favorite"`
	IsOnWatchList bool     `json:"is_on_my_watchlist"`
	CasualWins    int      `json:"casual_wins"`
	CasualLosses  int      `json:"casual_losses"`
	Cards         []Card   `json:"cards"`
	Houses        []string `json:"houses"`
	CardList      []string `json:"card_list"`
}

func LoadDeckFromFile(fileName string) (Deck, error) {
	deck := Deck{}

	bytes, e := ioutil.ReadFile(fileName)

	if e != nil {
		return deck, e
	}

	e = json.Unmarshal(bytes, &deck)

	if e != nil {
		return deck, e
	}

	return deck, nil
}
