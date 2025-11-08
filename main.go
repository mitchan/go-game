package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	cellSize = 32
)

type Sprite struct {
	Image *ebiten.Image
	X, Y  float64
}

type Game struct {
	Player *Sprite
}

func (game *Game) Update() error {
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

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Go game")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	playerImg, _, err := ebitenutil.NewImageFromFile("assets/images/player.png")
	if err != nil {
		log.Fatal(err)
	}

	player := &Sprite{
		Image: playerImg,
		X:     100,
		Y:     100,
	}

	if err := ebiten.RunGame(&Game{
		Player: player,
	}); err != nil {
		log.Fatal(err)
	}
}
