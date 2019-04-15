package keyforge

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	Running    bool
	Debug      bool
	Simulation bool
	Seed       int64
	Turn       int
	Round      int
	Players    []Player
	Bots       []*Bot
}

type BoardState struct {
}

func NewGame() *Game {
	game := new(Game)
	return game
}

func (g *Game) Start() {
	g.Running = true
	g.Debug = true
	g.Simulation = true

	if g.Seed == 0 {
		rand.Seed(time.Now().UTC().UnixNano())
	} else {
		rand.Seed(g.Seed)
	}

	deck, e := LoadDeckFromFile("../test/test_data/test_deck.json")

	if e != nil {
		fmt.Println(e)
		return
	}

	if g.Simulation {
		playerOne := NewBot()
		playerOne.Name = "Player one"
		playerOne.SetDeck(deck)
		playerOne.ShuffleDrawPile()

		playerTwo := NewBot()
		playerTwo.Name = "Player two"
		playerTwo.SetDeck(deck)
		playerTwo.ShuffleDrawPile()

		g.AddBot(playerOne)
		g.AddBot(playerTwo)
	}

	g.DetermineFirstPlayer()

	// draw up cards, determine mulligan
	if g.Simulation {
		for i := range g.Bots {
			bot := g.Bots[i]
			bot.DrawHand()

			if bot.FirstTurn {
				fmt.Println(bot.Name, "drawing an additional card for winning the toss.")
				bot.DrawCard()
			}

			if bot.DetermineMulligan() {
				fmt.Println(bot.Name, "chose to mulligan.")

				bot.SetDeck(deck)
				bot.ShuffleDrawPile()

				for i := 0; i < 5; i++ {
					bot.DrawCard()
				}

				if bot.FirstTurn {
					bot.DrawCard()
				}
			}

		}
	}

	for _, bot := range g.Bots {
		fmt.Println("Opening hand for", bot.Name)
		//fmt.Println(len(bot.HandPile))
		for _, card := range bot.HandPile {
			fmt.Println(card.CardTitle)
		}
	}

	for g.Running {
		g.Round = 1
		g.GameLoop()
		fmt.Println("#### Game results ####")
		fmt.Println("Round:", g.Round)
		fmt.Println("Turn:", g.Turn)
	}
}

func (g *Game) DetermineFirstPlayer() {
	if g.Simulation {
		roll := rand.Intn(len(g.Bots) - 1)
		bot := g.Bots[roll]
		bot.FirstTurn = true
		fmt.Println(bot.Name, "won the toss!")
	}
}

func (g *Game) RollDice() (int, int) {
	return rand.Intn(100), rand.Intn(100)
}

func (g *Game) AddPlayer(p *Player) {
	g.Players = append(g.Players, *p)
}

func (g *Game) AddBot(b *Bot) {
	g.Bots = append(g.Bots, b)
}

func (g *Game) GameLoop() {
	for g.Running {
		g.ExecuteTurn()
	}
}

func (g *Game) ExecuteRound() {

}

func (g *Game) ExecuteTurn() {
	var firstPlayer *Bot
	var secondPlayer *Bot
	g.Turn = 1

	for _, bot := range g.Bots {
		if bot.FirstTurn {
			firstPlayer = bot
		} else {
			secondPlayer = bot
		}
	}

	if firstPlayer.Amber > 6 {
		firstPlayer.ForgeKey()
	}

	if firstPlayer.Keys > 2 {
		fmt.Println("")
		fmt.Println("PLAYER ONE WINS THE GAME!")
		g.Running = false
		return
	}

	house := firstPlayer.DetermineActiveHouse()
	fmt.Println(firstPlayer.Name, "chose house", house)
	firstPlayer.PlayCards(house)
	firstPlayer.DrawHand()
	firstPlayer.PrettyPrintHand()

	g.Turn++

	if secondPlayer.Amber > 6 {
		secondPlayer.ForgeKey()
	}

	if secondPlayer.Keys > 2 {
		fmt.Println("")
		fmt.Println("PLAYER TWO WINS THE GAME!")
		g.Running = false
		return
	}

	house = secondPlayer.DetermineActiveHouse()
	fmt.Println(secondPlayer.Name, "chose house", house)
	secondPlayer.PlayCards(house)
	secondPlayer.DrawHand()
	secondPlayer.PrettyPrintHand()

	g.Round++
}
