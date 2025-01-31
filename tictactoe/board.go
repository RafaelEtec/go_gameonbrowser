package tictactoe

type Board struct {
	tiles [][]*Tile
}

func NewBoard() *Board {
	board := &Board{
		tiles: SetTiles(),
	}

	return board
}
