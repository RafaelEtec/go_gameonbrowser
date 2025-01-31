package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rafaeletec/go_gameonbrowser/tictactoe"
)

func main() {
	game := tictactoe.NewGame()

	ebiten.SetWindowSize(tictactoe.SCREEN_WIDTH*2, tictactoe.SCREEN_HEIGHT*2)
	ebiten.SetWindowTitle("TIC TAC TOE by Rafael Goulart")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
