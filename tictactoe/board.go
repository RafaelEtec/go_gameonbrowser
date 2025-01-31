package tictactoe

type Board struct {
	tiles [][]*Tile
}

func NewBoard(rows int, cols int) *Board {
	board := &Board{
		tiles: SetTiles(),
	}

	return board
}
