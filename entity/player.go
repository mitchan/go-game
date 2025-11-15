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

	speed    float64
	velocity math.Vector2D

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
		speed:    1.0,
		velocity: math.Vector2D{},
		Health:   100,
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

func (p *Player) Draw(screen *ebiten.Image, camera math.Vector2D) {
	opts := ebiten.DrawImageOptions{}

	activeAnimation := p.activeAnimation(p.velocity)
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

func (p *Player) Update(mapX, mapY int) {
	p.velocity.X = 0
	p.velocity.Y = 0

	// handle move
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.velocity.Y = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.velocity.X = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.velocity.Y = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.velocity.X = 1
	}

	newX := p.Sprite.X + p.velocity.X*p.speed
	newY := p.Sprite.Y + p.velocity.Y*p.speed

	if newX >= 0 && newX+constants.CellSize < float64(mapX) {
		p.Sprite.X = newX
	}
	if newY >= 0 && newY+constants.CellSize < float64(mapY) {
		p.Sprite.Y = newY
	}

	p.activeAnimation(p.velocity).Update()
}

func (p *Player) activeAnimation(velocity math.Vector2D) *animation.Animation {
	if velocity.X > 0 {
		return p.animations[right]
	}
	if velocity.X < 0 {
		return p.animations[left]
	}
	if velocity.Y > 0 {
		return p.animations[up]
	}
	if velocity.Y < 0 {
		return p.animations[down]
	}

	return p.animations[idle]
}
