// Package piece will handle game pieces.
package piece

import (
	"fmt"

	"github.com/jglouis/4eail40_2020/exercises/chess/model/coord"
	"github.com/jglouis/4eail40_2020/exercises/chess/model/player"
)

// Piece represents a game piece
type Piece interface {
	fmt.Stringer
	// Color returns the appartenance.
	Color() player.Color
	// Moves returns a set of valid move.
	Moves(isCapture bool) map[coord.ChessCoordinates]bool
	// NewPawn create a new pawn with a given color
	// Returns the pawn Piece
	NewPawn(color player.Color) Pawn
	// NewPawn create a new queen with a given color
	// Returns the queen Piece
	NewQueen(color player.Color) Queen
}
