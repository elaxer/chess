package chess

import "testing"

var edgePosition = PositionFromString("h8")

func TestNewSquares(t *testing.T) {
	squares := NewSquares(edgePosition)
	if len(squares.rows) != int(edgePosition.Rank) {
		t.Fatalf("expected 8 rows, got %d", len(squares.rows))
	}

	for _, row := range squares.rows {
		if len(row) != int(edgePosition.File) {
			t.Errorf("expected 8 squares, got %d", len(row))
		}
	}
}
