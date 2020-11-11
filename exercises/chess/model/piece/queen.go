package piece

import "github.com/jglouis/4eail40_2020/exercises/chess/model/player"

// Queen piece
type Queen struct {
	color player.Color
}

// NewQueen create a pawn with a given player color
func NewQueen(color player.Color) *Queen {
	return &Queen{color}
}
