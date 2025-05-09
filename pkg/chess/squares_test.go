package chess

import (
	"testing"

	"github.com/elaxer/chess/pkg/chess/position"
)

func TestNewSquares(t *testing.T) {
	squares := NewSquares(position.New(position.FileH, position.Rank8))
	if len(squares.squares) != 64 {
		t.Errorf("expected 8 squares, got %d", len(squares.squares))
	}
}
