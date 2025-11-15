package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/mitchan/go-game/constants"
	"github.com/mitchan/go-game/entity"
	"github.com/mitchan/go-game/math"
)

type Game struct {
	pigs      []*entity.Pig
	player    *entity.Player
	skeletons []*entity.Skeleton

	camera *Camera

	tilemapGrassImg *ebiten.Image
	tilemapJSON     *TilemapJSON
}

func NewGame() (*Game, error) {
	playerImg, _, err := ebitenutil.NewImageFromFile(constants.PlayerSpritePath)
	if err != nil {
		return nil, err
	}
	pigImg, _, err := ebitenutil.NewImageFromFile(constants.PigSpritePath)
	if err != nil {
		return nil, err
	}
	skeletonImg, _, err := ebitenutil.NewImageFromFile(constants.SkeletonSpritePath)
	if err != nil {
		return nil, err
	}
	tilemapGrass, _, err := ebitenutil.NewImageFromFile("assets/images/tileset-grass.png")
	if err != nil {
		return nil, err
	}

	tilemapJSON, err := NewTilemapJSON("assets/maps/map.json")

	return &Game{
		player: entity.NewPlayer(playerImg),
		skeletons: []*entity.Skeleton{
			entity.NewSkeleton(skeletonImg, 200.0, 200.0),
		},
		pigs: []*entity.Pig{
			entity.NewPig(pigImg, 32.0, 32.0),
			entity.NewPig(pigImg, 64.0, 64.0),
		},
		tilemapGrassImg: tilemapGrass,
		tilemapJSON:     tilemapJSON,
		camera:          NewCamera(0.0, 0.0),
	}, nil
}

func (g *Game) Update() error {
	g.player.Update(g.tilemapJSON.Width*constants.CellSize, g.tilemapJSON.Height*constants.CellSize)

	for _, enemy := range g.skeletons {
		enemy.Update(*g.player)
	}
	for _, pig := range g.pigs {
		pig.Update()
	}

	g.camera.FollowTarget(
		g.player.X+constants.CellSize/2,
		g.player.Y+constants.CellSize/2,
		constants.WindowWidth/constants.Zoom,
		constants.WindowHeight/constants.Zoom,
	)
	g.camera.Constrain(
		float64(g.tilemapJSON.Width)*constants.CellSize,
		float64(g.tilemapJSON.Height)*constants.CellSize,
		constants.WindowWidth/constants.Zoom,
		constants.WindowHeight/constants.Zoom,
	)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

	camera := math.Vector2D(*g.camera)

	g.tilemapJSON.Draw(screen, g.tilemapGrassImg, camera)

	for _, enemy := range g.skeletons {
		enemy.Draw(screen, camera)
	}
	for _, pig := range g.pigs {
		pig.Draw(screen, camera)
	}

	g.player.Draw(screen, camera)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return constants.WindowWidth / constants.Zoom, constants.WindowHeight / constants.Zoom
}
