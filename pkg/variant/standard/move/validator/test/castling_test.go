package validator_test

import (
	"testing"

	. "github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/move/move"
	"github.com/elaxer/chess/pkg/variant/standard/move/validator"
	"github.com/elaxer/chess/pkg/variant/standard/standardtest"
)

// todo добавить тесты с новым параметром Side
func TestValidateCastling(t *testing.T) {
	type fields struct {
		board Board
	}
	type args struct {
		castling move.Castling
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
				standardtest.MustNew(SideWhite, map[position.Position]Piece{
					position.FromString("e1"): standardtest.NewPiece("K"),
					position.FromString("g8"): standardtest.NewPiece("Q"),
					position.FromString("a1"): standardtest.NewPiece("R"),
					position.FromString("h1"): standardtest.NewPiece("R"),
					position.FromString("b6"): standardtest.NewPiece("r"),
				}),
			},
			args{move.CastlingShort},
			false,
		},
		{
			"long",
			fields{
				standardtest.MustNew(SideWhite, map[position.Position]Piece{
					position.FromString("e1"): standardtest.NewPiece("K"),
					position.FromString("a1"): standardtest.NewPiece("R"),
					position.FromString("h1"): standardtest.NewPiece("R"),
					position.FromString("g6"): standardtest.NewPiece("r"),
				}),
			},
			args{move.CastlingLong},
			false,
		},
		{
			"king_is_walked",
			fields{
				standardtest.MustNew(SideWhite, map[position.Position]Piece{
					position.FromString("e1"): standardtest.NewPieceW("K"),
					position.FromString("a1"): standardtest.NewPiece("R"),
					position.FromString("h1"): standardtest.NewPiece("R"),
				}),
			},
			args{move.CastlingShort},
			true,
		},
		{
			"rook_is_walked",
			fields{
				standardtest.MustNew(SideWhite, map[position.Position]Piece{
					position.FromString("e1"): standardtest.NewPiece("K"),
					position.FromString("a1"): standardtest.NewPiece("R"),
					position.FromString("h1"): standardtest.NewPieceW("R"),
				}),
			},
			args{move.CastlingShort},
			true,
		},
		{
			"let",
			fields{
				standardtest.MustNew(SideWhite, map[position.Position]Piece{
					position.FromString("e1"): standardtest.NewPiece("K"),
					position.FromString("a1"): standardtest.NewPiece("R"),
					position.FromString("g1"): standardtest.NewPiece("N"),
					position.FromString("h1"): standardtest.NewPiece("R"),
				}),
			},
			args{move.CastlingShort},
			true,
		},
		{
			"obstacle",
			fields{
				standardtest.MustNew(SideWhite, map[position.Position]Piece{
					position.FromString("e1"): standardtest.NewPiece("K"),
					position.FromString("a1"): standardtest.NewPiece("R"),
					position.FromString("h1"): standardtest.NewPiece("R"),
					position.FromString("g1"): standardtest.NewPiece("n"),
				}),
			},
			args{move.CastlingShort},
			true,
		},
		{
			"future_check",
			fields{
				standardtest.MustNew(SideWhite, map[position.Position]Piece{
					position.FromString("e1"): standardtest.NewPiece("K"),
					position.FromString("a1"): standardtest.NewPiece("R"),
					position.FromString("h1"): standardtest.NewPiece("R"),
					position.FromString("g8"): standardtest.NewPiece("r"),
				}),
			},
			args{move.CastlingShort},
			true,
		},
		{
			"attacked_castling_square",
			fields{
				standardtest.MustNew(SideWhite, map[position.Position]Piece{
					position.FromString("e1"): standardtest.NewPiece("K"),
					position.FromString("a1"): standardtest.NewPiece("R"),
					position.FromString("h1"): standardtest.NewPiece("R"),
					position.FromString("f8"): standardtest.NewPiece("r"),
				}),
			},
			args{move.CastlingShort},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateCastlingMove(tt.args.castling, tt.fields.board.Turn(), tt.fields.board, true)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateCastling() = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
