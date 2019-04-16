package keyforge

import "fmt"

// Player - This struct represents players within the game. This type is
// also the foundation of the Bot class, as it implements most if not all
// the functionality Player has.
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
	Artifacts   []Card
	Creatures   []Card
	FirstTurn   bool
	Amber       int
	Keys        int
	Chains      int
}

// NewPlayer - Returns a pointer to a new player object.
func NewPlayer() *Player {
	player := new(Player)
	player.DrawPile = make([]Card, 0)
	player.HandPile = make([]Card, 0)
	player.ArchivePile = make([]Card, 0)
	player.DiscardPile = make([]Card, 0)

	return player
}

// PrettyPrintHand - Debug function used to view the contents of a player's
// hand without making your eyes bleed.
func (p *Player) PrettyPrintHand() {
	fmt.Println()
	fmt.Println("########", p.Name, "'s hand ########")
	for _, card := range p.HandPile {
		fmt.Println(card.CardTitle)
	}
	fmt.Println()
}

// SetDeck - Sets a player's deck. If a deck has already been defined it will
// be cleared and replaced with the specified deck. This function is mainly
// useful for setting up players at the beginning of a game.
func (p *Player) SetDeck(deck Deck) {
	p.PlayerDeck = deck
	p.DrawPile = nil
	p.DrawPile = append(p.DrawPile, p.PlayerDeck.Cards...)
}

// ShuffleDrawPile - This function shuffles the player's draw pile (surprise).
func (p *Player) ShuffleDrawPile() {
	for i := 0; i < 10; i++ {
		p.DrawPile = Shuffle(p.DrawPile)
	}
}

// DrawCard - This function simulates a player drawing a card from the top
// of the draw pile into the player's hand. If the draw pile is found to be
// empty this function automatically shuffles the discard pile back into
// the draw pile.
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

// Discard - Discard a card from the player's hand. Cards discarded in
// this manner are sent to the player's discard pile.
func (p *Player) Discard(card Card) {
	p.HandPile = RemoveCard(p.HandPile, card)
	p.DiscardPile = AddCard(p.DiscardPile, card)
}

// ShuffleDiscardPile - This function transfers the contents of the discard
// pile to the draw pile and shuffles them. This is mostly useful for
// shuffling the discard into the draw pile after the player has exhausted
// the draw pile during play.
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

// DrawHand - Draws up a player hand to the appropriate number of cards.
// This function factors in handicaps from chains when drawing.
func (p *Player) DrawHand() {
	cardNumber := len(p.HandPile)

	if p.Debug {
		fmt.Println(cardNumber, "cards in hand, drawing", 6-cardNumber-p.CalculateChainHandicap(), "cards.")
	}

	// Draw back up to 6 cards, minus the handicap imposed by chains.
	for i := cardNumber; i < 6-p.CalculateChainHandicap(); i++ {
		p.DrawCard()
	}
}

// PlayCard - Play a card from the player's hand onto the board.
// TODO: Add logic to account for creatures and artifacts. Right now during
// simulated play they pretty much go directly to the discard pile rather
// than being used for reaping/amber gain/etc.
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

// ForgeKey - Attempt to forge a key given enough aember.
// TODO: Modify this function to account for board state. This function
// does not currently account for cards that either increase or decrease
// key forge costs.
func (p *Player) ForgeKey() bool {
	if p.Amber > 6 {
		fmt.Println(p.Name, "forges a key!")
		p.Keys++
		p.Amber -= 6
		return true
	}

	return false
}

// CalculateChainHandicap - Returns the total number of cards to reduce
// the player's hand upon drawing cards.
func (p *Player) CalculateChainHandicap() int {
	if p.Chains == 0 {
		return 0
	}

	chains := int(p.Chains / 6)

	if p.Chains < 6 {
		chains++
	}

	return chains
}

// DeployCreatureLeftFlank - This function places a creature card on the left
// flank of the battlefield
func (p *Player) DeployCreatureLeftFlank(card Card) []Card {
	return PrependCard(p.Creatures, card)
}

// DeployCreatureRightFlank - This function places a creature card on the right
// flank of the battlefield
func (p *Player) DeployCreatureRightFlank(card Card) []Card {
	return AddCard(p.Creatures, card)
}

// DeployCreature - Generic creature deploy function to be used in the event
// that the creatures array is empty. This function exists primarily for
// readability purposes since calling DeployCreatureRightFlank() in the event
// of an empty creature pile could be a bit confusing.
func (p *Player) DeployCreature(card Card) []Card {
	return AddCard(p.Creatures, card)
}
