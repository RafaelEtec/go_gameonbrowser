package tictactoe

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Tile struct {
	Img   *ebiten.Image
	Value int

	X int
	Y int

	Col int
	Row int
}

func SetTiles() [][]*Tile {
	url := "https://github.com/RafaelEtec/go_gameonbrowser/blob/6a1feebe7ebb17b5f8de29d1f93a445c19b2614b/tictactoe/assets/tiles.png?raw=true"
	img, err := ebitenutil.NewImageFromURL(url)
	if err != nil {
		log.Fatal(err)
	}

	tiles := [][]*Tile{
		{{}, {}, {}},
		{{}, {}, {}},
		{{}, {}, {}},
	}

	for c := 0; c < 3; c++ {
		for r := 0; r < 3; r++ {
			tile := tiles[c][r]

			tile.Img = img
			tile.Value = -1
			tile.Col = c
			tile.Row = r
			tile.X = c * 64
			tile.Y = r * 64
		}
	}

	return tiles
}
