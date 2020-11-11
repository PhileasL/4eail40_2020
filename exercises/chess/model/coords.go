// Package model contains the gameplay logic for the game of chess
package model

import "fmt"

// ChessCoordinates interface for flexible frame dimension or form choice
type ChessCoordinates interface {
	fmt.Stringer

	// Coord gets n'th coordinates component
	// Start with 0th coordinate
	// Returns an error if component doesn't exist
	Coord(n int) (int, error)
}
