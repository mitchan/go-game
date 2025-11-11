package spritesheet

import "image"

type SpriteSheet struct {
	WidthInTiles  int
	HeightInTiles int
	TileSize      int
}

func NewSpriteSheet(w, h, ts int) *SpriteSheet {
	return &SpriteSheet{
		WidthInTiles:  w,
		HeightInTiles: h,
		TileSize:      ts,
	}
}

func (s *SpriteSheet) Rect(index int) image.Rectangle {
	x := (index % s.WidthInTiles) * s.TileSize
	y := (index / s.WidthInTiles) * s.TileSize

	return image.Rect(x, y, x+s.TileSize, y+s.TileSize)
}
