package chess_test

import (
	"testing"

	. "github.com/elaxer/chess"
	"github.com/elaxer/chess/chesstest"
	"github.com/elaxer/chess/position"
)

var edgePosition = position.FromString("h8")

func TestSquares_FindByPosition(t *testing.T) {
	squares := NewSquares(edgePosition)
	king := chesstest.NewPiece("K")

	squares.PlacePiece(king, position.New(position.FileA, position.Rank4))
	p, err := squares.FindByPosition(position.FromString("a4"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if p != king {
		t.Errorf("expected %s, got %s", king, p)
	}
}

func TestSquares_MovePiece(t *testing.T) {
	king := chesstest.NewPiece("K")
	knight := chesstest.NewPiece("n")

	squares := NewSquares(edgePosition)
	squares.PlacePiece(king, position.FromString("e1"))
	squares.PlacePiece(knight, position.FromString("e2"))

	capturedPiece, err := squares.MovePiece(position.FromString("e1"), position.FromString("e2"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if capturedPiece == nil || capturedPiece != knight {
		t.Fatalf("expected %s, got %v", knight, capturedPiece)
	}
	if piece, _ := squares.FindByPosition(position.FromString("e1")); piece != nil {
		t.Errorf("expected nil, got %v", piece)
	}
	if piece, _ := squares.FindByPosition(position.FromString("e2")); piece != king {
		t.Errorf("expected %s, got %v", king, piece)
	}
}

func TestSquares_MovePieceTemporarily(t *testing.T) {
	queen := chesstest.NewPiece("Q")

	squares := NewSquares(edgePosition)
	squares.PlacePiece(queen, position.FromString("b4"))

	squares.MovePieceTemporarily(position.FromString("b4"), position.FromString("f8"), func() {
		if piece, _ := squares.FindByPosition(position.FromString("b4")); piece != nil {
			t.Errorf("expected nil, got %v", piece)
		}
		if piece, _ := squares.FindByPosition(position.FromString("f8")); piece != queen {
			t.Errorf("expected %s, got %v", queen, piece)
		}
	})

	if piece, _ := squares.FindByPosition(position.FromString("b4")); piece != queen {
		t.Errorf("expected %s, got %v", queen, piece)
	}
	if piece, _ := squares.FindByPosition(position.FromString("f8")); piece != nil {
		t.Errorf("expected nil, got %v", piece)
	}
}
