package mover_test

import (
	"testing"

	. "github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/variant/standard"
	"github.com/elaxer/chess/pkg/variant/standard/move"
	"github.com/elaxer/chess/pkg/variant/standard/mover"
	"github.com/elaxer/chess/pkg/variant/standard/piece"

	"github.com/elaxer/chess/pkg/chess/position"
)

func TestBoard_CastleShort(t *testing.T) {
	b := standard.NewFactory().CreateEmpty()
	squares := b.Squares()

	kingSquare := squares.GetByPosition(position.FromNotation("e1"))
	rookSquare := squares.GetByPosition(position.FromNotation("h1"))

	kingSquare.SetPiece(piece.NewKing(SideWhite))
	rookSquare.SetPiece(piece.NewRook(SideWhite))

	if _, err := new(mover.Castling).Make(move.CastlingShort, b); err != nil {
		t.Fatalf("castling failed: %v", err)
	}

	if !kingSquare.IsEmpty() {
		t.Fatalf("the king square should be empty")
	}
	if !rookSquare.IsEmpty() {
		t.Fatalf("the rook square should be empty")
	}

	kingSquareCastled := squares.GetByPosition(position.FromNotation("g1"))
	if kingSquareCastled.IsEmpty() {
		t.Fatalf("the king didn't castled")
	}

	rookSquareCastled := squares.GetByPosition(position.FromNotation("f1"))
	if rookSquareCastled.IsEmpty() {
		t.Errorf("the rook didn't castled")
	}
}

func TestBoard_CastleLong(t *testing.T) {
	b := standard.NewFactory().CreateEmpty()
	squares := b.Squares()

	kingSquare := squares.GetByPosition(position.FromNotation("e1"))
	rookSquare := squares.GetByPosition(position.FromNotation("a1"))

	kingSquare.SetPiece(piece.NewKing(SideWhite))
	rookSquare.SetPiece(piece.NewRook(SideWhite))

	if _, err := new(mover.Castling).Make(move.CastlingLong, b); err != nil {
		t.Fatalf("castling failed: %v", err)
	}

	if !kingSquare.IsEmpty() {
		t.Fatalf("kingSquare should be empty")
	}
	if !rookSquare.IsEmpty() {
		t.Fatalf("rookSquare should be empty")
	}

	if kingSquareCastled := squares.GetByPosition(position.FromNotation("c1")); kingSquareCastled.IsEmpty() {
		t.Fatalf("the king didn't castled")
	}
	if rookSquareCastled := squares.GetByPosition(position.FromNotation("d1")); rookSquareCastled.IsEmpty() {
		t.Errorf("the rook didn't castled")
	}
}

func TestBoard_CastleBlack(t *testing.T) {
	b := standard.NewFactory().CreateEmpty()
	b.NextTurn()

	squares := b.Squares()

	kingSquare := squares.GetByPosition(position.FromNotation("e8"))
	rookSquare := squares.GetByPosition(position.FromNotation("h8"))

	kingSquare.SetPiece(piece.NewKing(SideBlack))
	rookSquare.SetPiece(piece.NewRook(SideBlack))

	if _, err := new(mover.Castling).Make(move.CastlingShort, b); err != nil {
		t.Fatalf("castling failed: %v", err)
	}

	if !kingSquare.IsEmpty() {
		t.Fatalf("the king square should be empty")
	}
	if !rookSquare.IsEmpty() {
		t.Fatalf("the rook square should be empty")
	}

	kingSquareCastled := squares.GetByPosition(position.FromNotation("g8"))
	if kingSquareCastled.IsEmpty() {
		t.Fatalf("the king didn't castled")
	}

	rookSquareCastled := squares.GetByPosition(position.FromNotation("f8"))
	if rookSquareCastled.IsEmpty() {
		t.Errorf("the rook didn't castled")
	}
}
