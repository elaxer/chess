package chess_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/chess/chesstest"
)

var edgePosition = chess.PositionFromString("h8")

func TestSquares_FindByPosition(t *testing.T) {
	squares := chess.NewSquares(edgePosition)
	king := chesstest.NewPiece("K")

	if err := squares.PlacePiece(king, chess.NewPosition(chess.FileA, chess.Rank4)); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	p, err := squares.FindByPosition(chess.PositionFromString("a4"))
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

	squares := chess.NewSquares(edgePosition)
	if err := squares.PlacePiece(king, chess.PositionFromString("e1")); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := squares.PlacePiece(knight, chess.PositionFromString("e2")); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	capturedPiece, err := squares.MovePiece(
		chess.PositionFromString("e1"),
		chess.PositionFromString("e2"),
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if capturedPiece == nil || capturedPiece != knight {
		t.Fatalf("expected %s, got %v", knight, capturedPiece)
	}
	if piece, _ := squares.FindByPosition(chess.PositionFromString("e1")); piece != nil {
		t.Errorf("expected nil, got %v", piece)
	}
	if piece, _ := squares.FindByPosition(chess.PositionFromString("e2")); piece != king {
		t.Errorf("expected %s, got %v", king, piece)
	}
}

func TestSquares_MovePieceTemporarily(t *testing.T) {
	queen := chesstest.NewPiece("Q")

	squares := chess.NewSquares(edgePosition)
	if err := squares.PlacePiece(queen, chess.PositionFromString("b4")); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err := squares.MovePieceTemporarily(
		chess.PositionFromString("b4"),
		chess.PositionFromString("f8"),
		func() {
			if piece, _ := squares.FindByPosition(chess.PositionFromString("b4")); piece != nil {
				t.Errorf("expected nil, got %v", piece)
			}
			if piece, _ := squares.FindByPosition(chess.PositionFromString("f8")); piece != queen {
				t.Errorf("expected %s, got %v", queen, piece)
			}
		},
	)
	if err != nil {
		t.Fatalf("unexpected temporarily piece movement error: %v", err)
	}

	if piece, _ := squares.FindByPosition(chess.PositionFromString("b4")); piece != queen {
		t.Errorf("expected %s, got %v", queen, piece)
	}
	if piece, _ := squares.FindByPosition(chess.PositionFromString("f8")); piece != nil {
		t.Errorf("expected nil, got %v", piece)
	}
}
