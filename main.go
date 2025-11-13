package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mitchan/go-game/constants"
)

func main() {
	ebiten.SetWindowSize(constants.WindowWidth, constants.WindowHeight)
	ebiten.SetWindowTitle("Go game")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	game, err := NewGame()
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
