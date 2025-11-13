package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mitchan/go-game/animation"
	"github.com/mitchan/go-game/constants"
	"github.com/mitchan/go-game/math"
	"github.com/mitchan/go-game/spritesheet"
)

type Player struct {
	*Sprite
	Health float64

	dx int
	dy int

	// animation
	playerSpriteSheet *spritesheet.SpriteSheet
	animations        map[entityState]*animation.Animation
}

func NewPlayer(img *ebiten.Image) *Player {
	return &Player{
		Sprite: &Sprite{
			Image: img,
			X:     100,
			Y:     100,
		},
		dx:     0,
		dy:     0,
		Health: 100,
		playerSpriteSheet: &spritesheet.SpriteSheet{
			WidthInTiles:  6,
			HeightInTiles: 1,
			TileSize:      32,
		},
		animations: map[entityState]*animation.Animation{
			idle:  animation.NewAnimation(0, 5, 1, 5.0, false, false),
			up:    animation.NewAnimation(18, 23, 1, 5.0, false, false),
			right: animation.NewAnimation(24, 29, 1, 5.0, false, false),
			left:  animation.NewAnimation(24, 29, 1, 5.0, true, false),
			down:  animation.NewAnimation(30, 35, 1, 5.0, false, false),
		},
	}
}

func (p *Player) Draw(screen *ebiten.Image, camera math.Vector) {
	opts := ebiten.DrawImageOptions{}

	activeAnimation := p.activeAnimation(p.dx, p.dy)
	imageRect := p.playerSpriteSheet.Rect(activeAnimation.Frame())

	if activeAnimation.FlipH {
		opts.GeoM.Scale(-1, 1)
		opts.GeoM.Translate(constants.CellSize, 0)
	}

	opts.GeoM.Translate(p.X, p.Y)
	opts.GeoM.Translate(camera.X, camera.Y)

	screen.DrawImage(
		p.Image.SubImage(imageRect).(*ebiten.Image),
		&opts,
	)
}

func (p *Player) Update() {
	p.dx = 0
	p.dy = 0

	// handle move
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.dy = -2
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.dx = -2
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.dy = 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.dx = 2
	}

	p.Sprite.X += (float64)(p.dx)
	p.Sprite.Y += (float64)(p.dy)

	p.activeAnimation(p.dx, p.dy).Update()
}

func (p *Player) activeAnimation(dx, dy int) *animation.Animation {
	if dx > 0 {
		return p.animations[right]
	}
	if dx < 0 {
		return p.animations[left]
	}
	if dy > 0 {
		return p.animations[up]
	}
	if dy < 0 {
		return p.animations[down]
	}

	return p.animations[idle]
}
