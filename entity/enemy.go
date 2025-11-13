package entity

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mitchan/go-game/constants"
	"github.com/mitchan/go-game/math"
)

const (
	speed = 1
)

type Enemy struct {
	*Sprite
	Health        float64
	FollowsPlayer bool
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

func (e *Enemy) Draw(screen *ebiten.Image, camera math.Vector) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(e.X, e.Y)
	opts.GeoM.Translate(camera.X, camera.Y)
	screen.DrawImage(
		e.Image.SubImage(
			image.Rect(0, 0, constants.CellSize, constants.CellSize),
		).(*ebiten.Image),
		&opts,
	)
}

func (e *Enemy) Update(player Player) {
	if !e.FollowsPlayer {
		return
	}

	if player.X > e.X {
		e.X += speed
	} else if player.X < e.X {
		e.X -= speed
	}

	if player.Y > e.Y {
		e.Y += speed
	} else if player.Y < e.Y {
		e.Y -= speed
	}
}
