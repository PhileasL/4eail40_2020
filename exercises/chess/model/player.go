// Package model contains the gameplay logic for the game of chess
package model

// Color represents the color of a given player
// Either white or black
type Color int

const (
	// White is a Color
	White Color = iota
	// Black is a Color
	Black
)
