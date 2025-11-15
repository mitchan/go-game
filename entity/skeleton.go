package entity

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mitchan/go-game/constants"
	"github.com/mitchan/go-game/math"
)

type Skeleton struct {
	*Sprite
	health        float64
	followsPlayer bool
	speed         float64
}

func NewSkeleton(img *ebiten.Image, x, y float64) *Skeleton {
	return &Skeleton{
		Sprite: &Sprite{
			Image: img,
			X:     x,
			Y:     y,
		},
		health:        100,
		speed:         1.0,
		followsPlayer: true,
	}
}

func (e *Skeleton) Draw(screen *ebiten.Image, camera math.Vector) {
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

func (e *Skeleton) Update(player Player) {
	if !e.followsPlayer {
		return
	}

	if player.X > e.X {
		e.X += e.speed
	} else if player.X < e.X {
		e.X -= e.speed
	}

	if player.Y > e.Y {
		e.Y += e.speed
	} else if player.Y < e.Y {
		e.Y -= e.speed
	}
}
