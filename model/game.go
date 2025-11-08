package model

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	cellSize = 32
)

type Game struct {
	Player *Sprite
}

func (game *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		game.Player.Y -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		game.Player.X -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		game.Player.Y += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		game.Player.X += 2
	}
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

	// player
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(game.Player.X, game.Player.Y)
	screen.DrawImage(
		// grab a subimage of the spritesheet
		game.Player.Image.SubImage(
			image.Rect(0, 0, cellSize, cellSize),
		).(*ebiten.Image),
		&opts,
	)
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
