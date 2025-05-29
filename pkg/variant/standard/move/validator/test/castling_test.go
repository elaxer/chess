package validator_test

import (
	"testing"

	. "github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/variant/standard/move/move"
	"github.com/elaxer/chess/pkg/variant/standard/move/validator"
	"github.com/elaxer/chess/pkg/variant/standardtest"

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
				standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: standardtest.NewPiece("K", false), Position: FromNotation("e1")},
					{Piece: standardtest.NewPiece("Q", false), Position: FromNotation("g8")},
					{Piece: standardtest.NewPiece("R", false), Position: FromNotation("a1")},
					{Piece: standardtest.NewPiece("R", false), Position: FromNotation("h1")},
					{Piece: standardtest.NewPiece("r", false), Position: FromNotation("b6")},
				}),
			},
			args{move.CastlingShort},
			false,
		},
		{
			"long",
			fields{
				standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: standardtest.NewPiece("K", false), Position: FromNotation("e1")},
					{Piece: standardtest.NewPiece("R", false), Position: FromNotation("a1")},
					{Piece: standardtest.NewPiece("R", false), Position: FromNotation("h1")},
					{Piece: standardtest.NewPiece("r", false), Position: FromNotation("g6")},
				}),
			},
			args{move.CastlingLong},
			false,
		},
		{
			"king_is_walked",
			fields{
				standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: standardtest.NewPiece("K", true), Position: FromNotation("e1")},
					{Piece: standardtest.NewPiece("R", false), Position: FromNotation("a1")},
					{Piece: standardtest.NewPiece("R", false), Position: FromNotation("h1")},
				}),
			},
			args{move.CastlingShort},
			true,
		},
		{
			"rook_is_walked",
			fields{
				standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: standardtest.NewPiece("K", false), Position: FromNotation("e1")},
					{Piece: standardtest.NewPiece("R", false), Position: FromNotation("a1")},
					{Piece: standardtest.NewPiece("R", true), Position: FromNotation("h1")},
				}),
			},
			args{move.CastlingShort},
			true,
		},
		{
			"let",
			fields{
				standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: standardtest.NewPiece("K", false), Position: FromNotation("e1")},
					{Piece: standardtest.NewPiece("R", false), Position: FromNotation("a1")},
					{Piece: standardtest.NewPiece("N", false), Position: FromNotation("g1")},
					{Piece: standardtest.NewPiece("R", false), Position: FromNotation("h1")},
				}),
			},
			args{move.CastlingShort},
			true,
		},
		{
			"obstacle",
			fields{
				standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: standardtest.NewPiece("K", false), Position: FromNotation("e1")},
					{Piece: standardtest.NewPiece("R", false), Position: FromNotation("a1")},
					{Piece: standardtest.NewPiece("R", false), Position: FromNotation("h1")},
					{Piece: standardtest.NewPiece("n", false), Position: FromNotation("g1")},
				}),
			},
			args{move.CastlingShort},
			true,
		},
		{
			"future_check",
			fields{
				standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: standardtest.NewPiece("K", false), Position: FromNotation("e1")},
					{Piece: standardtest.NewPiece("R", false), Position: FromNotation("a1")},
					{Piece: standardtest.NewPiece("R", false), Position: FromNotation("h1")},
					{Piece: standardtest.NewPiece("r", false), Position: FromNotation("g8")},
				}),
			},
			args{move.CastlingShort},
			true,
		},
		{
			"attacked_castling_square",
			fields{
				standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: standardtest.NewPiece("K", false), Position: FromNotation("e1")},
					{Piece: standardtest.NewPiece("R", false), Position: FromNotation("a1")},
					{Piece: standardtest.NewPiece("R", false), Position: FromNotation("h1")},
					{Piece: standardtest.NewPiece("r", false), Position: FromNotation("f8")},
				}),
			},
			args{move.CastlingShort},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validator.ValidateCastling(tt.args.castling, tt.fields.board.Turn(), tt.fields.board, true); (err != nil) != tt.wantErr {
				t.Errorf("ValidateCastling() = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
