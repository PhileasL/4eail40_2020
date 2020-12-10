package board

import (
	"testing"

	"github.com/jglouis/4eail40_2020/exercises/chess/model/coord"

	"github.com/jglouis/4eail40_2020/exercises/chess/model/piece"
	"github.com/jglouis/4eail40_2020/exercises/chess/model/player"
)

type mockCoord int

// Coord returns x if n==0, y if n==1
func (c mockCoord) Coord(n int) (int, error) {
	return int(c), nil
}

func (c mockCoord) String() string {
	return "1"
}

func TestClassic_MovePiece(t *testing.T) {
	origin := coord.NewCartesian(0, 0)
	board := NewBoardClassic()
	pawn := piece.NewPawn(player.Black)
	queen := piece.NewQueen(player.Black)

	if err := board.PlacePieceAt(pawn, origin); err != nil {
		t.Error(err)
	}

	if err := board.PlacePieceAt(queen, origin.Right()); err != nil {
		t.Error(err)
	}

	if err := board.MovePiece(origin, origin.Left()); err != nil {
		t.Error("expected an error while moving left")
	}

	if err := board.MovePiece(origin, origin.UpRight()); err != nil {
		t.Error(err)
	}

	if board.PieceAt(origin.UpRight()) != pawn {
		t.Error("Could not find the moved piece at destination coordinates")
	}

	if err := board.MovePiece(origin.Right(), origin.UpRight()); err != nil {
		t.Error("expected the move to fail if destination was occupied")
	}

	if err := board.MovePiece(origin.Up(), origin.UpRight()); err != nil {
		t.Error("expected the move to fail if no piece at the origin of the move")
	}
}
