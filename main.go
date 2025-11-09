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

	player := entities.NewPlayer(playerImg)

	game := &Game{
		Player: player,
		Enemies: []*entities.Enemy{
			entities.NewEnemy(pigImg, 32.0, 32.0),
			entities.NewEnemy(pigImg, 64.0, 64.0),
		},
		tilemapGrassImg: tilemapGrass,
		tilemapJSON:     tilemapJSON,
		camera:          NewCamera(0.0, 0.0),
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
