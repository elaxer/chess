package formatter

import (
	"testing"

	. "github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	. "github.com/elaxer/chess/pkg/variant/standard"
	"github.com/elaxer/chess/pkg/variant/standard/move"
	. "github.com/elaxer/chess/pkg/variant/standard/piece"
)

func TestNormal_FormatSameFile(t *testing.T) {
	b := NewBoardFactory().CreateEmpty()
	squares := b.Squares()
	squares.AddPiece(NewQueen(SideWhite), position.FromNotation("d1"))
	squares.AddPiece(NewQueen(SideWhite), position.FromNotation("d8"))

	normalMove := &move.Normal{
		CheckMate:     &move.CheckMate{IsCheck: true},
		From:          position.FromNotation("d1"),
		To:            position.FromNotation("d4"),
		PieceNotation: NotationQueen,
		IsCapture:     true,
	}

	got, err := FormatNormalMove(normalMove, b)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	const expected = "Q1xd4+"
	if got != expected {
		t.Errorf("expected %s, got %s", expected, got)
	}
}

func TestNormal_FormatSameRank(t *testing.T) {
	b := NewBoardFactory().CreateEmpty()
	squares := b.Squares()
	squares.AddPiece(NewRook(SideBlack), position.FromNotation("a1"))
	squares.AddPiece(NewRook(SideBlack), position.FromNotation("g1"))

	normalMove := &move.Normal{
		CheckMate:     &move.CheckMate{IsMate: true},
		From:          position.FromNotation("a1"),
		To:            position.FromNotation("d1"),
		PieceNotation: NotationRook,
	}

	got, err := FormatNormalMove(normalMove, b)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	const expected = "Rad1#"
	if got != expected {
		t.Errorf("expected %s, got %s", expected, got)
	}
}

func TestNormal_FormatSameFileAndRank(t *testing.T) {
	b := NewBoardFactory().CreateEmpty()
	squares := b.Squares()
	squares.AddPiece(NewBishop(SideWhite), position.FromNotation("b7"))
	squares.AddPiece(NewBishop(SideWhite), position.FromNotation("f7"))
	squares.AddPiece(NewBishop(SideWhite), position.FromNotation("b3"))

	normalMove := &move.Normal{
		CheckMate:     &move.CheckMate{IsCheck: true},
		From:          position.FromNotation("b7"),
		To:            position.FromNotation("d5"),
		PieceNotation: NotationBishop,
	}

	got, err := FormatNormalMove(normalMove, b)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	const expected = "Bb7d5+"
	if got != expected {
		t.Errorf("expected %s, got %s", expected, got)
	}
}
