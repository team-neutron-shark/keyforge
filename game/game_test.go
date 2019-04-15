package keyforge

import (
	"testing"
)

func TestGameInit(t *testing.T) {
	game := NewGame()
	game.Start()
}
