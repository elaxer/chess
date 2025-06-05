package mover_test

import (
	"testing"

	. "github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/move/move"
	"github.com/elaxer/chess/pkg/variant/standard/move/mover"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
	"github.com/elaxer/chess/pkg/variant/standard/standardtest"
)

func TestPromotion_Make(t *testing.T) {
	board := standardtest.MustNew(SideWhite, map[position.Position]Piece{
		position.FromString("d7"): standardtest.NewPiece("P"),
		position.FromString("a1"): standardtest.NewPiece("K"),
		position.FromString("a8"): standardtest.NewPiece("k"),
	})

	promotion := move.NewPromotion(position.NewEmpty(), position.FromString("d8"), piece.NotationQueen)
	_, err := new(mover.Promotion).Make(promotion, board)
	if err != nil {
		t.Fatalf("promotion failed: %v", err)
	}

	queen, err := board.Squares().FindByPosition(position.FromString("d8"))
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
