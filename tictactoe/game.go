package tictactoe

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	SCREEN_WIDTH  = 192
	SCREEN_HEIGHT = 208

	FRAME_OX     = 0
	FRAME_OY     = 0
	FRAME_WIDTH  = 64
	FRAME_HEIGTH = 64

	ROWS = 3
	COLS = 3

	MESSAGE_WHOPLAYS  = "P%d turn"
	MESSAGE_MOVES     = "Moves: %d"
	MESSAGE_PLAYAGAIN = "P%d, move invalid!"
	MESSAGE_DRAW      = "It's a Draw!"
	MESSAGE_WINNER    = "P%d Wins!"
)

var (
	BACKGROUNDCOLOR = color.RGBA{0, 0, 0, 0}
)

type Game struct {
	message string
	moves   int

	player *Player
	board  *Board
}

func NewGame() *Game {
	firstToPlay := rand.Intn(2) + 1

	return &Game{
		message: fmt.Sprintf(MESSAGE_WHOPLAYS, firstToPlay),
		moves:   9,
		player:  NewPlayer(firstToPlay),
		board:   NewBoard(),
	}
}

func (g *Game) Update() error {

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}

	drawTiles(g, opts, screen)
	drawStats(g, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func drawTiles(g *Game, opts ebiten.DrawImageOptions, screen *ebiten.Image) {
	for col := 0; col < COLS; col++ {
		for row := 0; row < ROWS; row++ {
			tile := g.board.tiles[col][row]

			frameOffset := 1
			heightOffset := 1
			switch tile.Value {
			case 1:
				frameOffset = 64
				heightOffset = 2
			case 2:
				frameOffset = 128
				heightOffset = 3
			}

			opts.GeoM.Translate(float64(row)*FRAME_WIDTH, float64(col)*FRAME_HEIGTH)
			screen.DrawImage(
				tile.Img.SubImage(
					image.Rect(FRAME_OX, FRAME_OY+frameOffset, FRAME_WIDTH, FRAME_HEIGTH*heightOffset),
				).(*ebiten.Image),
				&opts,
			)
			opts.GeoM.Reset()
		}
	}
}

func drawStats(g *Game, screen *ebiten.Image) {
	moves := fmt.Sprintf(MESSAGE_MOVES, 9-g.moves)

	ebitenutil.DebugPrintAt(screen, g.message, 0, 192)
	ebitenutil.DebugPrintAt(screen, moves, 142, 192)
}
