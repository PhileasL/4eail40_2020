package piece

import (
	"fmt"

	"github.com/jglouis/4eail40_2020/exercises/chess/model/player"
)

// Queen piece
type Queen struct {
	color player.Color
}

// NewQueen create a pawn with a given player color
func NewQueen(color player.Color) *Queen {
	return &Queen{color}
}

//Color return the player.Color of the piece
func (q Queen) Color() player.Color {
	return q.color
}

func (q Queen) Moves(isCapture bool) {

}

func (q Queen) String() string {
	if q.color == player.White {
		return fmt.Sprintf("White")
	} else {
		return fmt.Sprintf("Black")
	}
}
