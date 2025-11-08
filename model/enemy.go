package model

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	*Sprite
	Health float64
}

func (enemy *Enemy) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(enemy.X, enemy.Y)
	screen.DrawImage(
		enemy.Image.SubImage(
			image.Rect(0, 0, cellSize, cellSize),
		).(*ebiten.Image),
		&opts,
	)
}
