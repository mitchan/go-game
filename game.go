package main

import (
	"image"
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

func (game *Game) Update() error {
	game.Player.Update()
	for _, enemy := range game.Enemies {
		enemy.Update(*game.Player)
	}
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

	game.tilemapJSON.Draw(screen, game.tilemapGrassImg)

	for _, enemy := range game.Enemies {
		enemy.Draw(screen)
	}

	// draw player
	// TODO: move to Player
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(game.Player.X, game.Player.Y)
	screen.DrawImage(
		game.Player.Image.SubImage(
			image.Rect(0, 0, constants.CellSize, constants.CellSize),
		).(*ebiten.Image),
		&opts,
	)
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return constants.WindowWidth, constants.WindowHeight
}
