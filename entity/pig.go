package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mitchan/go-game/animation"
	"github.com/mitchan/go-game/math"
	"github.com/mitchan/go-game/spritesheet"
)

type Pig struct {
	*Sprite
	Health float64

	// animation
	pigSpriteSheet *spritesheet.SpriteSheet
	animations     map[entityState]*animation.Animation
}

func NewPig(img *ebiten.Image, x, y float64) *Pig {
	return &Pig{
		Sprite: &Sprite{
			Image: img,
			X:     x,
			Y:     y,
		},
		pigSpriteSheet: &spritesheet.SpriteSheet{
			WidthInTiles:  2,
			HeightInTiles: 1,
			TileSize:      32,
		},
		animations: map[entityState]*animation.Animation{
			idle: animation.NewAnimation(0, 1, 1, 10.0, false, false),
		},
	}
}

func (p *Pig) Update() {
	p.animations[idle].Update()
}

func (p *Pig) Draw(screen *ebiten.Image, camera math.Vector) {
	opts := ebiten.DrawImageOptions{}

	imageRect := p.pigSpriteSheet.Rect(p.animations[idle].Frame())

	opts.GeoM.Translate(p.X, p.Y)
	opts.GeoM.Translate(camera.X, camera.Y)

	screen.DrawImage(
		p.Image.SubImage(imageRect).(*ebiten.Image),
		&opts,
	)
}
