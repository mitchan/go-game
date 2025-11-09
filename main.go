package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/mitchan/go-game/constants"
	"github.com/mitchan/go-game/entities"
)

func main() {
	ebiten.SetWindowSize(constants.WindowWidth, constants.WindowHeight)
	ebiten.SetWindowTitle("Go game")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	playerImg, _, err := ebitenutil.NewImageFromFile(constants.PlayerSpritePath)
	if err != nil {
		log.Fatal(err)
	}
	pigImg, _, err := ebitenutil.NewImageFromFile(constants.PigSpritePath)
	if err != nil {
		log.Fatal(err)
	}
	tilemapGrass, _, err := ebitenutil.NewImageFromFile("assets/images/tileset-grass.png")
	if err != nil {
		// handle error
		log.Fatal(err)
	}

	tilemapJSON, err := NewTilemapJSON("assets/maps/map.json")

	player := &entities.Player{
		Sprite: &entities.Sprite{
			Image: playerImg,
			X:     100,
			Y:     100,
		},
		Health: 100,
	}

	game := &Game{
		Player: player,
		Enemies: []*entities.Enemy{
			{
				Sprite: &entities.Sprite{
					Image: pigImg,
					X:     64,
					Y:     64,
				},
				Health: 100,
			},
			{
				Sprite: &entities.Sprite{
					Image: pigImg,
					X:     164,
					Y:     164,
				},
				Health: 100,
			},
		},
		tilemapGrassImg: tilemapGrass,
		tilemapJSON:     tilemapJSON,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
