package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mitchan/go-game/constants"
	"github.com/mitchan/go-game/entities"
	"github.com/mitchan/go-game/math"
)

type Game struct {
	Player          *entities.Player
	Enemies         []*entities.Enemy
	tilemapJSON     *TilemapJSON
	tilemapGrassImg *ebiten.Image
	camera          *Camera
}

func (g *Game) Update() error {
	g.Player.Update()
	for _, enemy := range g.Enemies {
		enemy.Update(*g.Player)
	}

	g.camera.FollowTarget(
		g.Player.X+constants.CellSize/2,
		g.Player.Y+constants.CellSize/2,
		constants.WindowWidth/constants.Zoom,
		constants.WindowHeight/constants.Zoom,
	)
	g.camera.Constrain(
		float64(g.tilemapJSON.Layers[0].Width)*constants.CellSize,
		float64(g.tilemapJSON.Layers[0].Height)*constants.CellSize,
		constants.WindowWidth/constants.Zoom,
		constants.WindowHeight/constants.Zoom,
	)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

	camera := math.Vector(*g.camera)

	g.tilemapJSON.Draw(screen, g.tilemapGrassImg, camera)

	for _, enemy := range g.Enemies {
		enemy.Draw(screen, camera)
	}

	g.Player.Draw(screen, camera)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return constants.WindowWidth / constants.Zoom, constants.WindowHeight / constants.Zoom
}
