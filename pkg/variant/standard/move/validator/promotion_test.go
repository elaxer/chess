package validator_test

import (
	"testing"

	. "github.com/elaxer/chess/pkg/chess"
	. "github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard"
	"github.com/elaxer/chess/pkg/variant/standard/move"
	"github.com/elaxer/chess/pkg/variant/standard/move/validator"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
)

func TestPromotion_Validate(t *testing.T) {
	b := standard.NewBoardFactory().CreateEmpty()
	squares := b.Squares()

	squares.AddPiece(piece.NewPawn(SideWhite), FromNotation("d7"))
	squares.AddPiece(piece.NewKing(SideWhite), FromNotation("a1"))
	squares.AddPiece(piece.NewPawn(SideBlack), FromNotation("a8"))

	promotionMv := &move.Promotion{CheckMate: new(move.CheckMate), To: FromNotation("d8"), NewPiece: NotationQueen}
	if err := validator.ValidatePromotion(promotionMv, b); err != nil {
		t.Errorf("validation failed: %v", err)
	}
}
