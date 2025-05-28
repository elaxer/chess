package validator_test

import (
	"testing"

	. "github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/variant/standard/move"
	"github.com/elaxer/chess/pkg/variant/standard/move/validator"
	"github.com/elaxer/chess/pkg/variant/standarttest"

	. "github.com/elaxer/chess/pkg/chess/position"
)

// todo добавить тесты с новым параметром Side
func TestValidateCastling(t *testing.T) {
	type fields struct {
		board Board
	}
	type args struct {
		castling move.CastlingType
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"short",
			fields{
				standarttest.NewEmpty(SideWhite, []standarttest.Placement{
					{Piece: standarttest.NewPiece("K", SideWhite, false), Position: FromNotation("e1")},
					{Piece: standarttest.NewPiece("Q", SideWhite, false), Position: FromNotation("g8")},
					{Piece: standarttest.NewPiece("R", SideWhite, false), Position: FromNotation("a1")},
					{Piece: standarttest.NewPiece("R", SideWhite, false), Position: FromNotation("h1")},
					{Piece: standarttest.NewPiece("R", SideBlack, false), Position: FromNotation("b6")},
				}),
			},
			args{move.CastlingShort},
			false,
		},
		{
			"long",
			fields{
				standarttest.NewEmpty(SideWhite, []standarttest.Placement{
					{Piece: standarttest.NewPiece("K", SideWhite, false), Position: FromNotation("e1")},
					{Piece: standarttest.NewPiece("R", SideWhite, false), Position: FromNotation("a1")},
					{Piece: standarttest.NewPiece("R", SideWhite, false), Position: FromNotation("h1")},
					{Piece: standarttest.NewPiece("R", SideBlack, false), Position: FromNotation("g6")},
				}),
			},
			args{move.CastlingLong},
			false,
		},
		{
			"king_is_walked",
			fields{
				standarttest.NewEmpty(SideWhite, []standarttest.Placement{
					{Piece: standarttest.NewPiece("K", SideWhite, true), Position: FromNotation("e1")},
					{Piece: standarttest.NewPiece("R", SideWhite, false), Position: FromNotation("a1")},
					{Piece: standarttest.NewPiece("R", SideWhite, false), Position: FromNotation("h1")},
				}),
			},
			args{move.CastlingShort},
			true,
		},
		{
			"rook_is_walked",
			fields{
				standarttest.NewEmpty(SideWhite, []standarttest.Placement{
					{Piece: standarttest.NewPiece("K", SideWhite, false), Position: FromNotation("e1")},
					{Piece: standarttest.NewPiece("R", SideWhite, false), Position: FromNotation("a1")},
					{Piece: standarttest.NewPiece("R", SideWhite, true), Position: FromNotation("h1")},
				}),
			},
			args{move.CastlingShort},
			true,
		},
		{
			"let",
			fields{
				standarttest.NewEmpty(SideWhite, []standarttest.Placement{
					{Piece: standarttest.NewPiece("K", SideWhite, false), Position: FromNotation("e1")},
					{Piece: standarttest.NewPiece("R", SideWhite, false), Position: FromNotation("a1")},
					{Piece: standarttest.NewPiece("N", SideWhite, false), Position: FromNotation("g1")},
					{Piece: standarttest.NewPiece("R", SideWhite, false), Position: FromNotation("h1")},
				}),
			},
			args{move.CastlingShort},
			true,
		},
		{
			"obstacle",
			fields{
				standarttest.NewEmpty(SideWhite, []standarttest.Placement{
					{Piece: standarttest.NewPiece("K", SideWhite, false), Position: FromNotation("e1")},
					{Piece: standarttest.NewPiece("R", SideWhite, false), Position: FromNotation("a1")},
					{Piece: standarttest.NewPiece("R", SideWhite, false), Position: FromNotation("h1")},
					{Piece: standarttest.NewPiece("N", SideBlack, false), Position: FromNotation("g1")},
				}),
			},
			args{move.CastlingShort},
			true,
		},
		{
			"future_check",
			fields{
				standarttest.NewEmpty(SideWhite, []standarttest.Placement{
					{Piece: standarttest.NewPiece("K", SideWhite, false), Position: FromNotation("e1")},
					{Piece: standarttest.NewPiece("R", SideWhite, false), Position: FromNotation("a1")},
					{Piece: standarttest.NewPiece("R", SideWhite, false), Position: FromNotation("h1")},
					{Piece: standarttest.NewPiece("R", SideBlack, false), Position: FromNotation("g8")},
				}),
			},
			args{move.CastlingShort},
			true,
		},
		{
			"attacked_castling_square",
			fields{
				standarttest.NewEmpty(SideWhite, []standarttest.Placement{
					{Piece: standarttest.NewPiece("K", SideWhite, false), Position: FromNotation("e1")},
					{Piece: standarttest.NewPiece("R", SideWhite, false), Position: FromNotation("a1")},
					{Piece: standarttest.NewPiece("R", SideWhite, false), Position: FromNotation("h1")},
					{Piece: standarttest.NewPiece("R", SideBlack, false), Position: FromNotation("f8")},
				}),
			},
			args{move.CastlingShort},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validator.ValidateCastling(tt.args.castling, tt.fields.board.Turn(), tt.fields.board); (err != nil) != tt.wantErr {
				t.Errorf("ValidateCastling() = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
