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
	squares.AddPiece(NewQueen(SideWhite), position.New(position.FileD, position.Rank1))
	squares.AddPiece(NewQueen(SideWhite), position.New(position.FileD, position.Rank8))

	normalMove := &move.Normal{
		CheckMate:     &move.CheckMate{IsCheck: true},
		From:          position.New(position.FileD, position.Rank1),
		To:            position.New(position.FileD, position.Rank4),
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
	squares.AddPiece(NewRook(SideBlack), position.New(position.FileA, position.Rank1))
	squares.AddPiece(NewRook(SideBlack), position.New(position.FileG, position.Rank1))

	normalMove := &move.Normal{
		CheckMate:     &move.CheckMate{IsMate: true},
		From:          position.New(position.FileA, position.Rank1),
		To:            position.New(position.FileD, position.Rank1),
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
	squares.AddPiece(NewBishop(SideWhite), position.New(position.FileB, position.Rank7))
	squares.AddPiece(NewBishop(SideWhite), position.New(position.FileF, position.Rank7))
	squares.AddPiece(NewBishop(SideWhite), position.New(position.FileB, position.Rank3))

	normalMove := &move.Normal{
		CheckMate:     &move.CheckMate{IsCheck: true},
		From:          position.New(position.FileB, position.Rank7),
		To:            position.New(position.FileD, position.Rank5),
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
