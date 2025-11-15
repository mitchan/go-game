package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mitchan/go-game/animation"
	"github.com/mitchan/go-game/constants"
	"github.com/mitchan/go-game/math"
	"github.com/mitchan/go-game/spritesheet"
)

type Skeleton struct {
	*Sprite
	health        float64
	followsPlayer bool
	speed         float64
	velocity      math.Vector2D

	// animation
	playerSpriteSheet *spritesheet.SpriteSheet
	animations        map[entityState]*animation.Animation
}

func NewSkeleton(img *ebiten.Image, x, y float64) *Skeleton {
	return &Skeleton{
		Sprite: &Sprite{
			Image: img,
			X:     x,
			Y:     y,
		},
		health: 100,
		speed:  0.5,
		velocity: math.Vector2D{
			X: 0,
			Y: 0,
		},
		followsPlayer: true,
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

func (s *Skeleton) Draw(screen *ebiten.Image, camera math.Vector2D) {
	opts := ebiten.DrawImageOptions{}

	activeAnimation := s.activeAnimation(s.velocity)
	imageRect := s.playerSpriteSheet.Rect(activeAnimation.Frame())

	if activeAnimation.FlipH {
		opts.GeoM.Scale(-1, 1)
		opts.GeoM.Translate(constants.CellSize, 0)
	}

	opts.GeoM.Translate(s.X, s.Y)
	opts.GeoM.Translate(camera.X, camera.Y)
	screen.DrawImage(s.Image.SubImage(imageRect).(*ebiten.Image), &opts)
}

func (s *Skeleton) Update(player Player) {
	s.velocity.X = 0
	s.velocity.Y = 0

	if !s.followsPlayer {
		return
	}

	if player.X > s.X {
		s.velocity.X = 1
	} else if player.X < s.X {
		s.velocity.X = -1
	}

	if player.Y > s.Y {
		s.velocity.Y = 1
	} else if player.Y < s.Y {
		s.velocity.Y = -1
	}

	s.Sprite.X += s.velocity.X * s.speed
	s.Sprite.Y += s.velocity.Y * s.speed

	s.activeAnimation(s.velocity).Update()
}

func (s *Skeleton) activeAnimation(velocity math.Vector2D) *animation.Animation {
	if velocity.X > 0 {
		return s.animations[right]
	}
	if velocity.X < 0 {
		return s.animations[left]
	}
	if velocity.Y > 0 {
		return s.animations[up]
	}
	if velocity.Y < 0 {
		return s.animations[down]
	}

	return s.animations[idle]
}
