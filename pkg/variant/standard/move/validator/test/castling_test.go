package validator_test

import (
	"testing"

	. "github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/move/move"
	"github.com/elaxer/chess/pkg/variant/standard/move/validator"
	"github.com/elaxer/chess/pkg/variant/standardtest"
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
				standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: standardtest.NewPiece("K"), Position: position.FromString("e1")},
					{Piece: standardtest.NewPiece("Q"), Position: position.FromString("g8")},
					{Piece: standardtest.NewPiece("R"), Position: position.FromString("a1")},
					{Piece: standardtest.NewPiece("R"), Position: position.FromString("h1")},
					{Piece: standardtest.NewPiece("r"), Position: position.FromString("b6")},
				}),
			},
			args{move.CastlingShort},
			false,
		},
		{
			"long",
			fields{
				standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: standardtest.NewPiece("K"), Position: position.FromString("e1")},
					{Piece: standardtest.NewPiece("R"), Position: position.FromString("a1")},
					{Piece: standardtest.NewPiece("R"), Position: position.FromString("h1")},
					{Piece: standardtest.NewPiece("r"), Position: position.FromString("g6")},
				}),
			},
			args{move.CastlingLong},
			false,
		},
		{
			"king_is_walked",
			fields{
				standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: standardtest.NewPieceW("K"), Position: position.FromString("e1")},
					{Piece: standardtest.NewPiece("R"), Position: position.FromString("a1")},
					{Piece: standardtest.NewPiece("R"), Position: position.FromString("h1")},
				}),
			},
			args{move.CastlingShort},
			true,
		},
		{
			"rook_is_walked",
			fields{
				standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: standardtest.NewPiece("K"), Position: position.FromString("e1")},
					{Piece: standardtest.NewPiece("R"), Position: position.FromString("a1")},
					{Piece: standardtest.NewPieceW("R"), Position: position.FromString("h1")},
				}),
			},
			args{move.CastlingShort},
			true,
		},
		{
			"let",
			fields{
				standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: standardtest.NewPiece("K"), Position: position.FromString("e1")},
					{Piece: standardtest.NewPiece("R"), Position: position.FromString("a1")},
					{Piece: standardtest.NewPiece("N"), Position: position.FromString("g1")},
					{Piece: standardtest.NewPiece("R"), Position: position.FromString("h1")},
				}),
			},
			args{move.CastlingShort},
			true,
		},
		{
			"obstacle",
			fields{
				standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: standardtest.NewPiece("K"), Position: position.FromString("e1")},
					{Piece: standardtest.NewPiece("R"), Position: position.FromString("a1")},
					{Piece: standardtest.NewPiece("R"), Position: position.FromString("h1")},
					{Piece: standardtest.NewPiece("n"), Position: position.FromString("g1")},
				}),
			},
			args{move.CastlingShort},
			true,
		},
		{
			"future_check",
			fields{
				standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: standardtest.NewPiece("K"), Position: position.FromString("e1")},
					{Piece: standardtest.NewPiece("R"), Position: position.FromString("a1")},
					{Piece: standardtest.NewPiece("R"), Position: position.FromString("h1")},
					{Piece: standardtest.NewPiece("r"), Position: position.FromString("g8")},
				}),
			},
			args{move.CastlingShort},
			true,
		},
		{
			"attacked_castling_square",
			fields{
				standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: standardtest.NewPiece("K"), Position: position.FromString("e1")},
					{Piece: standardtest.NewPiece("R"), Position: position.FromString("a1")},
					{Piece: standardtest.NewPiece("R"), Position: position.FromString("h1")},
					{Piece: standardtest.NewPiece("r"), Position: position.FromString("f8")},
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
