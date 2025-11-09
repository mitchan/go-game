package entities

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mitchan/go-game/constants"
)

const (
	speed = 1
)

type Enemy struct {
	*Sprite
	Health float64
}

func NewEnemy(img *ebiten.Image, x, y float64) *Enemy {
	return &Enemy{
		Sprite: &Sprite{
			Image: img,
			X:     x,
			Y:     y,
		},
		Health: 100,
	}
}

func (enemy *Enemy) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(enemy.X, enemy.Y)
	screen.DrawImage(
		enemy.Image.SubImage(
			image.Rect(0, 0, constants.CellSize, constants.CellSize),
		).(*ebiten.Image),
		&opts,
	)
}

func (enemy *Enemy) Update(player Player) {
	if player.X > enemy.X {
		enemy.X += speed
	} else if player.X < enemy.X {
		enemy.X -= speed
	}

	if player.Y > enemy.Y {
		enemy.Y += speed
	} else if player.Y < enemy.Y {
		enemy.Y -= speed
	}
}
