package mover_test

import (
	"testing"

	. "github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/board"
	"github.com/elaxer/chess/pkg/variant/standard/move/move"
	"github.com/elaxer/chess/pkg/variant/standard/move/mover"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
)

func TestPromotion_Make(t *testing.T) {
	b := board.NewFactory().CreateEmpty(SideWhite)
	squares := b.Squares()

	squares.PlacePiece(piece.NewPawn(SideWhite), position.FromString("d7"))
	squares.PlacePiece(piece.NewKing(SideWhite), position.FromString("a1"))
	squares.PlacePiece(piece.NewKing(SideBlack), position.FromString("a8"))

	promotion := &move.Promotion{
		Normal:           &move.Normal{To: position.FromString("d8")},
		NewPieceNotation: piece.NotationQueen,
	}
	_, err := new(mover.Promotion).Make(promotion, b)
	if err != nil {
		t.Fatalf("promotion failed: %v", err)
	}

	queen, err := squares.GetByPosition(position.FromString("d8"))
	if err != nil {
		t.Fatal(err)
	}
	if queen == nil {
		t.Fatalf("the queen didn't appear on the board")
	}
	if queen.Notation() != piece.NotationQueen {
		t.Errorf("the piece should be a queen")
	}
}
