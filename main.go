package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/mitchan/go-game/model"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Go game")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	playerImg, _, err := ebitenutil.NewImageFromFile("assets/images/player.png")
	if err != nil {
		log.Fatal(err)
	}

	player := &model.Player{
		Sprite: &model.Sprite{
			Image: playerImg,
			X:     100,
			Y:     100,
		},
		Health: 100,
	}

	if err := ebiten.RunGame(&model.Game{
		Player: player,
	}); err != nil {
		log.Fatal(err)
	}
}
