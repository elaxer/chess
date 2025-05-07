package resolver_test

import (
	"testing"

	. "github.com/elaxer/chess/pkg/chess"
	. "github.com/elaxer/chess/pkg/chess/position"
	. "github.com/elaxer/chess/pkg/variant/standard"
	"github.com/elaxer/chess/pkg/variant/standard/move/resolver"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
)

func TestNormal_CreateMove(t *testing.T) {
	b := NewBoardFactory().CreateEmpty()
	squares := b.Squares()

	squares.AddPiece(piece.NewPawn(SideWhite), FromNotation("a2"))
	squares.AddPiece(piece.NewPawn(SideWhite), FromNotation("b2"))
	squares.AddPiece(piece.NewPawn(SideWhite), FromNotation("c2"))
	squares.AddPiece(piece.NewPawn(SideWhite), FromNotation("d2"))
	squares.AddPiece(piece.NewPawn(SideWhite), FromNotation("e2"))
	squares.AddPiece(piece.NewPawn(SideWhite), FromNotation("f2"))
	squares.AddPiece(piece.NewPawn(SideWhite), FromNotation("g2"))
	squares.AddPiece(piece.NewPawn(SideWhite), FromNotation("h2"))

	from, err := resolver.ResolveFrom(Position{}, FromNotation("e4"), NotationPawn, b)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expected := FromNotation("e2")
	if from != expected {
		t.Errorf("expected from position %v, got %v", expected, from)
	}
}

func TestNormal_CreateMoveSameFile(t *testing.T) {
	b := NewBoardFactory().CreateEmpty()
	b.NextTurn()
	squares := b.Squares()

	squares.AddPiece(piece.NewRook(SideBlack), FromNotation("a8"))
	squares.AddPiece(piece.NewRook(SideBlack), FromNotation("f8"))

	from, err := resolver.ResolveFrom(Position{File: FileA}, FromNotation("b8"), NotationRook, b)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expected := FromNotation("a8")
	if from != expected {
		t.Errorf("expected from position %v, got %v", expected, from)
	}
}

func TestNormal_CreateMoveSameRank(t *testing.T) {
	b := NewBoardFactory().CreateEmpty()
	squares := b.Squares()

	squares.AddPiece(piece.NewRook(SideWhite), FromNotation("a1"))
	squares.AddPiece(piece.NewRook(SideWhite), FromNotation("a8"))

	from, err := resolver.ResolveFrom(Position{Rank: Rank1}, FromNotation("a5"), NotationRook, b)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expected := FromNotation("a1")
	if from != expected {
		t.Errorf("expected from position %v, got %v", expected, from)
	}
}

func TestNormal_CreateMoveFullFrom(t *testing.T) {
	b := NewBoardFactory().CreateEmpty()
	squares := b.Squares()

	squares.AddPiece(piece.NewBishop(SideBlack), FromNotation("b2"))
	squares.AddPiece(piece.NewBishop(SideBlack), FromNotation("f2"))
	squares.AddPiece(piece.NewBishop(SideBlack), FromNotation("b6"))

	from, err := resolver.ResolveFrom(FromNotation("f2"), FromNotation("d4"), NotationBishop, b)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expected := FromNotation("f2")
	if from != expected {
		t.Errorf("expected from position %v, got %v", expected, from)
	}
}

func TestNormal_CreateMoveKnights(t *testing.T) {
	b := NewBoardFactory().CreateEmpty()
	squares := b.Squares()

	squares.AddPiece(piece.NewKnight(SideWhite), FromNotation("g1"))
	squares.AddPiece(piece.NewKnight(SideWhite), FromNotation("c3"))

	from, err := resolver.ResolveFrom(Position{File: FileG}, FromNotation("e2"), NotationKnight, b)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expected := FromNotation("g1")
	if from != expected {
		t.Errorf("expected from position %v, got %v", expected, from)
	}
}
