package entities

import "github.com/hajimehoshi/ebiten/v2"

type Player struct {
	*Sprite
	Health float64
}

func NewPlayer(img *ebiten.Image) *Player {
	return &Player{
		Sprite: &Sprite{
			Image: img,
			X:     100,
			Y:     100,
		},
		Health: 100,
	}
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
}
