package piece

import (
	"fmt"

	"github.com/jglouis/4eail40_2020/exercises/chess/model/player"
)

// Pawn piece
type Pawn struct {
	color player.Color
}

// NewPawn create a pawn with a given player color
func NewPawn(color player.Color) *Pawn {
	return &Pawn{color}
}

//Color return the player.Color of the piece
func (p Pawn) Color() player.Color {
	return p.color
}

func (p Pawn) Moves(isCapture bool) {

}

func (p Pawn) String() string {
	if p.color == player.White {
		return fmt.Sprintf("White")
	} else {
		return fmt.Sprintf("Black")
	}
}
