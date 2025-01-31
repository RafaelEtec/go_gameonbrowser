package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rafaeletec/go_gameonbrowser/tictactoe"
)

func main() {
	game := tictactoe.NewGame()

	ebiten.SetWindowSize(tictactoe.ScreenWidth, tictactoe.ScreenHeight)
	ebiten.SetWindowTitle("TicTacToe by Rafael Goulart")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
