package chess

import (
	"testing"

	"github.com/elaxer/chess/pkg/chess/position"
)

type mockPiece struct {
	notation string
	side     Side
}

func (m *mockPiece) Side() Side {
	return m.side
}

func (m *mockPiece) IsMoved() bool {
	return false
}

func (m *mockPiece) MarkMoved() {
}

func (m *mockPiece) PseudoMoves(from position.Position, squares *Squares) position.Set {
	return nil
}

func (m *mockPiece) Notation() string {
	return m.notation
}

func (m *mockPiece) Weight() uint8 {
	return 0
}

func (m *mockPiece) String() string {
	return m.side.String() + m.notation
}

func TestNewSquares(t *testing.T) {
	squares := NewSquares(position.New(position.FileH, position.Rank8))
	if len(squares.rows) != 8 {
		t.Errorf("expected 8 rows, got %d", len(squares.rows))
	}
	for _, row := range squares.rows {
		if len(squares.rows) != 8 {
			t.Errorf("expected 8 squares, got %d", len(row))
		}
	}
}

func TestSquares_GetByPosition(t *testing.T) {
	squares := NewSquares(position.New(position.FileH, position.Rank8))
	king := &mockPiece{"K", SideWhite}

	squares.rows[3][0] = king
	p, err := squares.GetByPosition(position.FromNotation("a4"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if p != king {
		t.Errorf("expected %s, got %s", king, p)
	}
}

func TestSquares_MovePiece(t *testing.T) {
	king := &mockPiece{"K", SideWhite}
	knight := &mockPiece{"N", SideBlack}

	squares := NewSquares(position.New(position.FileH, position.Rank8))
	squares.PlacePiece(king, position.FromNotation("e1"))
	squares.PlacePiece(knight, position.FromNotation("e2"))

	capturedPiece, err := squares.MovePiece(position.FromNotation("e1"), position.FromNotation("e2"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if capturedPiece == nil || capturedPiece != knight {
		t.Fatalf("expected %s, got %v", knight, capturedPiece)
	}
	if piece, _ := squares.GetByPosition(position.FromNotation("e1")); piece != nil {
		t.Errorf("expected nil, got %v", piece)
	}
	if piece, _ := squares.GetByPosition(position.FromNotation("e2")); piece != king {
		t.Errorf("expected %s, got %v", king, piece)
	}
}

func TestSquares_MovePieceTemporarily(t *testing.T) {
	queen := &mockPiece{"Q", SideWhite}

	squares := NewSquares(position.New(position.FileH, position.Rank8))
	squares.PlacePiece(queen, position.FromNotation("b4"))

	squares.MovePieceTemporarily(position.FromNotation("b4"), position.FromNotation("f8"), func() {
		if piece, _ := squares.GetByPosition(position.FromNotation("b4")); piece != nil {
			t.Errorf("expected nil, got %v", piece)
		}
		if piece, _ := squares.GetByPosition(position.FromNotation("f8")); piece != queen {
			t.Errorf("expected %s, got %v", queen, piece)
		}
	})

	if piece, _ := squares.GetByPosition(position.FromNotation("b4")); piece != queen {
		t.Errorf("expected %s, got %v", queen, piece)
	}
	if piece, _ := squares.GetByPosition(position.FromNotation("f8")); piece != nil {
		t.Errorf("expected nil, got %v", piece)
	}
}
