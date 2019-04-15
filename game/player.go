package keyforge

import "fmt"

type PlayerEntity interface {
	GetDeck() Deck
	GetDrawPile() []Card
	GetDiscardPile() []Card
	GetArchivePile() []Card
	GetPurgePile() []Card
}

type Player struct {
	Name        string
	Game        *Game
	Debug       bool
	PlayerDeck  Deck
	HandPile    []Card
	DrawPile    []Card
	DiscardPile []Card
	ArchivePile []Card
	PurgePile   []Card
	FirstTurn   bool
	Amber       int
	Keys        int
}

func NewPlayer() *Player {
	player := new(Player)
	player.DrawPile = make([]Card, 0)
	player.HandPile = make([]Card, 0)
	player.ArchivePile = make([]Card, 0)
	player.DiscardPile = make([]Card, 0)

	return player
}

func (p *Player) PrettyPrintHand() {
	fmt.Println()
	fmt.Println("########", p.Name, "'s hand ########")
	for _, card := range p.HandPile {
		fmt.Println(card.CardTitle)
	}
	fmt.Println()
}

func (p Player) GetDeck() Deck {
	return p.PlayerDeck
}

func (p Player) GetDrawPile() []Card {
	return p.DrawPile
}

func (p Player) GetDiscardPile() []Card {
	return p.DiscardPile
}

func (p Player) GetArchivePile() []Card {
	return p.ArchivePile
}

func (p Player) GetPurgePile() []Card {
	return p.PurgePile
}

func SetDeck(player PlayerEntity, deck Deck) {

}

func (p *Player) SetDeck(deck Deck) {
	p.PlayerDeck = deck
	p.DrawPile = nil
	p.DrawPile = append(p.DrawPile, p.PlayerDeck.Cards...)
}

func (p *Player) ShuffleDrawPile() {
	for i := 0; i < 10; i++ {
		p.DrawPile = Shuffle(p.DrawPile)
	}
}

func (p *Player) DrawCard() {
	count := len(p.DrawPile)
	index := len(p.DrawPile) - 1

	if index < 0 {
		index = 0
	}

	if count == 0 {
		fmt.Println("Draw pile empty, shuffling into discard.")
		p.ShuffleDiscardPile()
	}

	card := p.DrawPile[index]
	p.DrawPile = PopCard(p.DrawPile)
	p.HandPile = AddCard(p.HandPile, card)
}

func (p *Player) Discard(card Card) {
	p.HandPile = RemoveCard(p.HandPile, card)
	p.DiscardPile = AddCard(p.DiscardPile, card)
}
func (p *Player) ShuffleDiscardPile() {
	p.DrawPile = append(p.DrawPile, p.DiscardPile...)
	p.DiscardPile = nil

	for _, card := range p.DiscardPile {
		p.DiscardPile = RemoveCard(p.DrawPile, card)
		p.DrawPile = AddCard(p.DrawPile, card)
	}
	for i := 0; i < 10; i++ {
		p.DrawPile = Shuffle(p.DrawPile)
	}
}

func (p *Player) DrawHand() {
	cardNumber := len(p.HandPile)

	if p.Debug {
		fmt.Println(cardNumber, "cards in hand, drawing", 6-cardNumber, "cards.")
	}
	for i := cardNumber; i < 6; i++ {
		p.DrawCard()
	}
}

func (p *Player) PlayCard(card Card) {
	foundCard, e := FindCardByID(p.HandPile, card.ID)

	if e != nil {
		fmt.Println(e)
		return
	}

	p.HandPile = RemoveCard(p.HandPile, foundCard)
	p.DiscardPile = AddCard(p.DiscardPile, foundCard)

	fmt.Println(p.Name, "played card", foundCard.CardTitle)

	if card.Amber > 0 {
		fmt.Println(p.Name, "gains", card.Amber, "amber.")
		p.Amber += card.Amber
	}

}

func (p *Player) ForgeKey() {
	if p.Amber > 6 {
		fmt.Println(p.Name, "forges a key!")
		p.Keys++
		p.Amber -= 6
	}
}
