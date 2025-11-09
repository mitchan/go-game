package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mitchan/go-game/constants"
	"github.com/mitchan/go-game/entities"
)

type Game struct {
	Player          *entities.Player
	Enemies         []*entities.Enemy
	tilemapJSON     *TilemapJSON
	tilemapGrassImg *ebiten.Image
}

func (g *Game) Update() error {
	g.Player.Update()
	for _, enemy := range g.Enemies {
		enemy.Update(*g.Player)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

	g.tilemapJSON.Draw(screen, g.tilemapGrassImg)

	for _, enemy := range g.Enemies {
		enemy.Draw(screen)
	}

	g.Player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return constants.WindowWidth / constants.Zoom, constants.WindowHeight / constants.Zoom
}
