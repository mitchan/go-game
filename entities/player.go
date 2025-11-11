package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mitchan/go-game/animation"
	"github.com/mitchan/go-game/math"
	"github.com/mitchan/go-game/spritesheet"
)

type Player struct {
	*Sprite
	Health float64

	// animation
	playerSpriteSheet *spritesheet.SpriteSheet
	idleAnimation     *animation.Animation
}

func NewPlayer(img *ebiten.Image) *Player {
	return &Player{
		Sprite: &Sprite{
			Image: img,
			X:     100,
			Y:     100,
		},
		Health: 100,
		playerSpriteSheet: &spritesheet.SpriteSheet{
			WidthInTiles:  6,
			HeightInTiles: 1,
			TileSize:      32,
		},
		idleAnimation: animation.NewAnimation(0, 5, 1, 5.0),
	}
}

func (p *Player) Draw(screen *ebiten.Image, camera math.Vector) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(p.X, p.Y)
	opts.GeoM.Translate(camera.X, camera.Y)
	screen.DrawImage(
		p.Image.SubImage(
			p.playerSpriteSheet.Rect(p.idleAnimation.Frame()),
		).(*ebiten.Image),
		&opts,
	)
}

func (p *Player) Update() {
	// handle move
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.Y -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.X -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.Y += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.X += 2
	}

	p.idleAnimation.Update()
}
