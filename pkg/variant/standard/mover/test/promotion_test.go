package mover_test

import (
	"testing"

	. "github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard"
	"github.com/elaxer/chess/pkg/variant/standard/move"
	"github.com/elaxer/chess/pkg/variant/standard/mover"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
)

func TestPromotion_Make(t *testing.T) {
	b := standard.NewFactory().CreateEmpty(SideWhite)
	squares := b.Squares()

	squares.AddPiece(piece.NewPawn(SideWhite), position.FromNotation("d7"))
	squares.AddPiece(piece.NewKing(SideWhite), position.FromNotation("a1"))
	squares.AddPiece(piece.NewKing(SideBlack), position.FromNotation("a8"))

	promotion := &move.Promotion{
		Normal:   &move.Normal{CheckMate: new(move.CheckMate), To: position.FromNotation("d8")},
		NewPiece: NotationQueen,
	}
	_, err := new(mover.Promotion).Make(promotion, b)
	if err != nil {
		t.Fatalf("promotion failed: %v", err)
	}

	queen := squares.GetByPosition(position.FromNotation("d8"))
	if queen.IsEmpty() {
		t.Fatalf("the queen didn't appear on the board")
	}
	if queen.Piece.Notation() != NotationQueen {
		t.Errorf("the piece should be a queen")
	}
}
