package tictactoe

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 600
	ScreenHeight = 600
	boardRows    = 3
	boardCols    = 3
)

var (
	backgroundColor = color.RGBA{0, 0, 0, 0}
)

type Game struct {
	board *Board
}

func NewGame() *Game {
	return &Game{
		board: NewBoard(boardRows, boardCols),
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
