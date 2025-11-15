package main

import (
	"encoding/json"
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mitchan/go-game/constants"
	"github.com/mitchan/go-game/math"
)

type TilemapLayerJSON struct {
	Data []int `json:"data"`

	Width  int `json:"width"`
	Height int `json:"height"`
}

type TilemapJSON struct {
	Layers []TilemapLayerJSON `json:"layers"`

	Width  int `json:"width"`
	Height int `json:"height"`
}

func NewTilemapJSON(filepath string) (*TilemapJSON, error) {
	contents, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var tilemapJSON TilemapJSON
	err = json.Unmarshal(contents, &tilemapJSON)
	if err != nil {
		return nil, err
	}

	return &tilemapJSON, nil
}

func (t *TilemapJSON) Draw(
	screen *ebiten.Image,
	tilemapImg *ebiten.Image,
	camera math.Vector2D,
) {
	opts := ebiten.DrawImageOptions{}

	for _, layer := range t.Layers {
		for index, id := range layer.Data {
			x := index % layer.Width
			y := index / layer.Width

			x *= constants.CellSize
			y *= constants.CellSize

			srcX := (id - 1) % 10
			srcY := (id - 1) / 10

			srcX *= constants.CellSize
			srcY *= constants.CellSize

			opts.GeoM.Translate(float64(x), float64(y))
			opts.GeoM.Translate(camera.X, camera.Y)

			screen.DrawImage(
				tilemapImg.SubImage(image.Rect(srcX, srcY, srcX+constants.CellSize, srcY+constants.CellSize)).(*ebiten.Image),
				&opts,
			)

			// reset the opts for the next tile
			opts.GeoM.Reset()
		}
	}
}
