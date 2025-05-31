package validator_test

import (
	"testing"

	. "github.com/elaxer/chess/pkg/chess"
	. "github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/move/move"
	"github.com/elaxer/chess/pkg/variant/standard/move/validator"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
	"github.com/elaxer/chess/pkg/variant/standardtest"
)

func TestValidatePromotion(t *testing.T) {
	type args struct {
		move  *move.Promotion
		board Board
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid",
			args{
				&move.Promotion{
					Normal:           &move.Normal{To: FromString("d8")},
					NewPieceNotation: piece.NotationQueen,
				},
				standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: piece.NewPawn(SideWhite), Position: FromString("d7")},
				}),
			},
			false,
		},
		{
			"concurrent_pawns",
			args{
				&move.Promotion{
					Normal:           &move.Normal{From: Position{File: FileA}, To: FromString("b1")},
					NewPieceNotation: piece.NotationQueen,
				},
				standardtest.NewEmpty(SideBlack, []standardtest.Placement{
					{Piece: piece.NewPawn(SideBlack), Position: FromString("a2")},
					{Piece: piece.NewPawn(SideBlack), Position: FromString("c2")},
					{Piece: piece.NewKnight(SideWhite), Position: FromString("b1")},
				}),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			standardtest.ResolvePromotion(tt.args.move, tt.args.board)

			err := validator.ValidatePromotion(tt.args.move, tt.args.board)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidatePromotion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
