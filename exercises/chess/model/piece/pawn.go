package piece

import "github.com/jglouis/4eail40_2020/exercises/chess/model/player"

// Pawn piece
type Pawn struct {
	color player.Color
}

// NewPawn create a pawn with a given player color
func NewPawn(color player.Color) Pawn {
	return Pawn{color}
}
