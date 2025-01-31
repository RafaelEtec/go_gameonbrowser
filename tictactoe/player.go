package tictactoe

type Player struct {
	IsPlaying int
}

func NewPlayer(firstToPlay int) *Player {
	player := &Player{
		IsPlaying: firstToPlay,
	}

	return player
}

func Next(whosPlaying int) int {
	switch whosPlaying {
	case 1:
		return 2
	}
	return 1
}
