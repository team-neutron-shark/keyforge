package keyforgevault

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	keyforge "keyforge/game"
	"net/http"
	"net/url"
	"strconv"
)

type Credentials struct {
	IDToken     string `json:"id_token"`
	AccessToken string `json:"access_token"`
	Nonce       string `json:"nonce"`
	APIToken    string `json:"-"`
}

type VaultUser struct {
	ID                   string `json:"id"`
	UserName             string `json:"username"`
	Email                string `json:"email"`
	Avatar               string `json:"avatar"`
	Location             string `json:"location"`
	QRCode               string `json:"qr_code"`
	AccountURL           string `json:"account_url"`
	Language             string `json:"language"`
	AmberShards          int    `json:"amber_shards"`
	AmberShardsCollected int    `json:"amber_shards_collected"`
	Keys                 int    `json:"keys"`
	Token                string `json:"-"`
}

type PartialDeckJSON struct {
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
	CardList      []string `json:"cards"`
}

type FullDeckJSON struct {
	PartialDeckJSON
	Links FullDeckLinkJSON `json:"_links"`
}

type FullDeckLinkJSON struct {
	Houses   []string `json:"houses"`
	CardList []string `json:"cards"`
}

type SearchVaultUser struct {
	User  VaultUser `json:"user"`
	Token string    `json:"token"`
}

type SearchVaultUserJSON struct {
	Data SearchVaultUser `json:"data"`
}

type PartialDeckSearchJSON struct {
	Count int               `json:"count"`
	Decks []PartialDeckJSON `json:"data"`
}

type SearchDeckLink struct {
	Houses []string `json:"houses"`
}

type RetrieveDeckJSON struct {
	Deck   FullDeckJSON `json:"data"`
	Linked DeckLinkJSON `json:"_linked"`
}

type DeckLinkJSON struct {
	Houses []HouseJSON     `json:"houses"`
	Cards  []keyforge.Card `json:"cards"`
	Notes  []string        `json:"notes"`
}

type HouseJSON struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type DeckQuery struct {
	Page          int
	PageSize      int
	Query         string
	MinimumLevel  int
	MaximumLevel  int
	MinimumChains int
	MaximumChains int
	Ordering      string
	Values        url.Values
}

func (d *DeckQuery) GetQueryString() (string, error) {
	if d.Values == nil {
		d.Values = url.Values{}
	}

	if d.Page == 0 {
		d.Page = 1
	}

	if d.PageSize == 0 {
		d.PageSize = 1
	}

	page := strconv.Itoa(d.Page)
	pageSize := strconv.Itoa(d.PageSize)

	d.Values.Add("page", page)
	d.Values.Add("page_size", pageSize)
	d.Values.Add("search", d.Query)

	if d.MinimumChains > 0 && d.MaximumChains > 0 && d.MaximumChains >= d.MinimumChains {
		chainString := fmt.Sprintf("%d,%d", d.MinimumChains, d.MaximumChains)
		d.Values.Add("chains", chainString)
	}

	if d.MinimumLevel > 0 && d.MaximumLevel > 0 && d.MaximumLevel >= d.MinimumLevel {
		levelString := fmt.Sprintf("%d,%d", d.MinimumLevel, d.MaximumLevel)
		d.Values.Add("power_level", levelString)
	}

	if len(d.Ordering) > 0 {
		d.Values.Add("ordering", d.Ordering)
	}

	return d.Values.Encode(), nil
}

func Login(userName, password string) (VaultUser, error) {
	credentials := Credentials{}

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	params := url.Values{}
	params.Add("display", "popup")
	params.Add("scope", "openid profile email")
	params.Add("response_type", "id_token token")
	params.Add("client_id", "keyforge-web-portal")
	params.Add("state", "/")
	params.Add("redirect_uri", "https://www.keyforgegame.com/authorize")
	params.Add("nonce", "4HM~Z5f8")

	loginForm := params
	loginForm.Add("login", userName)
	loginForm.Add("password", password)

	path := fmt.Sprintf("https://account.asmodee.net/en/signin?%s", params.Encode())

	response, e := client.PostForm(path, loginForm)

	if e != nil {
		return VaultUser{}, e
	}

	redirectLocation, e := response.Location()

	if e != nil {
		return VaultUser{}, e
	}

	redirectLocation.RawQuery = redirectLocation.Fragment
	redirectParams := redirectLocation.Query()

	credentials.AccessToken = redirectParams.Get("access_token")
	credentials.IDToken = redirectParams.Get("id_token")
	credentials.Nonce = "4HM~Z5f8"

	jsonObject, e := json.Marshal(&credentials)

	if e != nil {
		return VaultUser{}, e
	}

	path = fmt.Sprintf("https://www.keyforgegame.com/api/users/login/asmodee/")

	contentLength := strconv.Itoa(len(jsonObject))
	request, e := http.NewRequest("POST", path, bytes.NewBuffer(jsonObject))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Content-Length", contentLength)

	if e != nil {
		return VaultUser{}, e
	}

	response, e = client.Do(request)

	if e != nil {
		return VaultUser{}, e
	}

	body, e := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	userResult := SearchVaultUserJSON{}
	json.Unmarshal(body, &userResult)

	userResult.Data.User.Token = userResult.Data.Token

	return userResult.Data.User, nil
}

func SearchDecks(vaultUser *VaultUser, deckQuery *DeckQuery) (PartialDeckSearchJSON, error) {
	client := &http.Client{}
	authHeader := fmt.Sprintf("Token %s", vaultUser.Token)
	params, e := deckQuery.GetQueryString()

	if e != nil {
		return PartialDeckSearchJSON{}, e
	}

	path := fmt.Sprintf("https://www.keyforgegame.com/api/decks/?%s", params)

	request, e := http.NewRequest("GET", path, nil)
	request.Header.Add("Authorization", authHeader)

	if e != nil {
		return PartialDeckSearchJSON{}, e
	}

	response, e := client.Do(request)

	body, e := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	result := PartialDeckSearchJSON{}
	e = json.Unmarshal(body, &result)

	if e != nil {
		return PartialDeckSearchJSON{}, e
	}

	return result, nil
}

func RetrieveDeck(vaultUser *VaultUser, deckID string) (keyforge.Deck, error) {
	newDeck := keyforge.Deck{}
	client := &http.Client{}
	deckJSON := RetrieveDeckJSON{}
	authHeader := fmt.Sprintf("Token %s", vaultUser.Token)
	path := fmt.Sprintf("https://www.keyforgegame.com/api/decks/%s/?links=cards,notes", deckID)

	request, e := http.NewRequest("GET", path, nil)
	request.Header.Add("Authorization", authHeader)

	if e != nil {
		return keyforge.Deck{}, e
	}

	response, e := client.Do(request)

	body, e := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	e = json.Unmarshal(body, &deckJSON)

	if e != nil {
		return keyforge.Deck{}, e
	}

	newDeck.CasualLosses = deckJSON.Deck.CasualLosses
	newDeck.CasualWins = deckJSON.Deck.CasualWins
	newDeck.Chains = deckJSON.Deck.Chains
	newDeck.Expansion = deckJSON.Deck.Expansion
	//newDeck.Houses = deckJSON.Deck.Houses
	newDeck.ID = deckJSON.Deck.ID
	newDeck.IsMyDeck = deckJSON.Deck.IsMyDeck
	newDeck.IsMyFavorite = deckJSON.Deck.IsMyFavorite
	newDeck.IsOnWatchList = deckJSON.Deck.IsOnWatchList
	newDeck.Losses = deckJSON.Deck.Losses
	newDeck.Name = deckJSON.Deck.Name
	newDeck.Notes = deckJSON.Deck.Notes
	newDeck.Wins = deckJSON.Deck.Wins
	newDeck.CardList = deckJSON.Deck.Links.CardList

	for _, cardID := range newDeck.CardList {
		for _, card := range deckJSON.Linked.Cards {
			if card.ID == cardID {
				newDeck.Cards = append(newDeck.Cards, card)
				break
			}
		}
	}

	return newDeck, nil
}
