package validator_test

import (
	"testing"

	. "github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/variant/standard"
	"github.com/elaxer/chess/pkg/variant/standard/move"
	"github.com/elaxer/chess/pkg/variant/standard/move/validator"
	"github.com/elaxer/chess/pkg/variant/standard/piece"

	. "github.com/elaxer/chess/pkg/chess/position"
)

func TestCastling_Validate(t *testing.T) {
	type Square struct {
		Piece    Piece
		Position Position
		IsMoved  bool
	}
	type fields struct {
		Turn   Side
		Pieces []*Square
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
		// {
		// 	"short",
		// 	fields{
		// 		SideWhite,
		// 		[]*Square{
		// 			{piece.NewKing(SideWhite), FromNotation("e1"), false},
		// 			{piece.NewQueen(SideWhite), FromNotation("g8"), false},
		// 			{piece.NewRook(SideWhite), FromNotation("a1"), false},
		// 			{piece.NewRook(SideWhite), FromNotation("h1"), false},
		// 			{piece.NewRook(SideBlack), FromNotation("b6"), false},
		// 		},
		// 	},
		// 	args{move.CastlingShort},
		// 	false,
		// },
		{
			"long",
			fields{
				SideWhite,
				[]*Square{
					{piece.NewKing(SideWhite), FromNotation("e1"), false},
					{piece.NewRook(SideWhite), FromNotation("a1"), false},
					{piece.NewRook(SideWhite), FromNotation("h1"), false},
					{piece.NewRook(SideBlack), FromNotation("g6"), false},
				},
			},
			args{move.CastlingLong},
			false,
		},
		// {
		// 	"king_is_walked",
		// 	fields{
		// 		SideWhite,
		// 		[]*Square{
		// 			{piece.NewKing(SideWhite), FromNotation("e1"), true},
		// 			{piece.NewRook(SideWhite), FromNotation("a1"), false},
		// 			{piece.NewRook(SideWhite), FromNotation("h1"), false},
		// 		},
		// 	},
		// 	args{move.CastlingShort},
		// 	true,
		// },
		// {
		// 	"rook_is_walked",
		// 	fields{
		// 		SideWhite,
		// 		[]*Square{
		// 			{piece.NewKing(SideWhite), FromNotation("e1"), false},
		// 			{piece.NewRook(SideWhite), FromNotation("a1"), false},
		// 			{piece.NewRook(SideWhite), FromNotation("h1"), true},
		// 		},
		// 	},
		// 	args{move.CastlingShort},
		// 	true,
		// },
		// {
		// 	"let",
		// 	fields{
		// 		SideWhite,
		// 		[]*Square{
		// 			{piece.NewKing(SideWhite), FromNotation("e1"), false},
		// 			{piece.NewRook(SideWhite), FromNotation("a1"), false},
		// 			{piece.NewKnight(SideWhite), FromNotation("g1"), false},
		// 			{piece.NewRook(SideWhite), FromNotation("h1"), false},
		// 		},
		// 	},
		// 	args{move.CastlingShort},
		// 	true,
		// },
		// {
		// 	"obstacle",
		// 	fields{
		// 		SideWhite,
		// 		[]*Square{
		// 			{piece.NewKing(SideWhite), FromNotation("e1"), false},
		// 			{piece.NewRook(SideWhite), FromNotation("a1"), false},
		// 			{piece.NewRook(SideWhite), FromNotation("h1"), false},
		// 			{piece.NewKnight(SideBlack), FromNotation("g1"), false},
		// 		},
		// 	},
		// 	args{move.CastlingShort},
		// 	true,
		// },
		// {
		// 	"future_check",
		// 	fields{
		// 		SideWhite,
		// 		[]*Square{
		// 			{piece.NewKing(SideWhite), FromNotation("e1"), false},
		// 			{piece.NewRook(SideWhite), FromNotation("a1"), false},
		// 			{piece.NewRook(SideWhite), FromNotation("h1"), false},
		// 			{piece.NewRook(SideBlack), FromNotation("g8"), false},
		// 		},
		// 	},
		// 	args{move.CastlingShort},
		// 	true,
		// },
		// {
		// 	"attacked_castling_square",
		// 	fields{
		// 		SideWhite,
		// 		[]*Square{
		// 			{piece.NewKing(SideWhite), FromNotation("e1"), false},
		// 			{piece.NewRook(SideWhite), FromNotation("a1"), false},
		// 			{piece.NewRook(SideWhite), FromNotation("h1"), false},
		// 			{piece.NewRook(SideBlack), FromNotation("f8"), false},
		// 		},
		// 	},
		// 	args{move.CastlingShort},
		// 	true,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := standard.NewBoardFactory().CreateEmpty()
			if tt.fields.Turn == SideBlack {
				b.NextTurn()
			}
			for _, square := range tt.fields.Pieces {
				if square.IsMoved {
					square.Piece.SetMoved()
				}

				b.Squares().AddPiece(square.Piece, square.Position)
			}

			if err := validator.ValidateCastling(tt.args.castling, b); (err != nil) != tt.wantErr {
				t.Errorf("Castling.ValidateMove() = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCastling_ValidateShort(t *testing.T) {
	b := standard.NewBoardFactory().CreateEmpty()
	squares := b.Squares()

	squares.AddPiece(piece.NewKing(SideWhite), FromNotation("e1"))
	squares.AddPiece(piece.NewRook(SideWhite), FromNotation("a1"))
	squares.AddPiece(piece.NewRook(SideWhite), FromNotation("h1"))

	if err := validator.ValidateCastling(move.CastlingShort, b); err != nil {
		t.Errorf("the white should be able castle")
	}
}

func TestCastling_ValidateLong(t *testing.T) {
	b := standard.NewBoardFactory().CreateEmpty()
	squares := b.Squares()

	squares.AddPiece(piece.NewKing(SideWhite), FromNotation("e1"))
	squares.AddPiece(piece.NewRook(SideWhite), FromNotation("a1"))
	squares.AddPiece(piece.NewRook(SideWhite), FromNotation("h1"))

	if err := validator.ValidateCastling(move.CastlingLong, b); err != nil {
		t.Errorf("the white should be able to castle")
	}
}

func TestCastling_ValidateKingWalked(t *testing.T) {
	b := standard.NewBoardFactory().CreateEmpty()
	squares := b.Squares()

	king := piece.NewKing(SideWhite)
	king.SetMoved()

	squares.AddPiece(king, FromNotation("e1"))
	squares.AddPiece(piece.NewRook(SideWhite), FromNotation("a1"))
	squares.AddPiece(piece.NewRook(SideWhite), FromNotation("h1"))

	if err := validator.ValidateCastling(move.CastlingShort, b); err == nil {
		t.Errorf("the white shouldn't be able castle, got error: %v", err)
	}
}

func TestCastling_ValidateRookWalked(t *testing.T) {
	b := standard.NewBoardFactory().CreateEmpty()
	squares := b.Squares()

	rook := piece.NewRook(SideWhite)
	rook.SetMoved()

	squares.AddPiece(piece.NewKing(SideWhite), FromNotation("e1"))
	squares.AddPiece(piece.NewRook(SideWhite), FromNotation("a1"))
	squares.AddPiece(rook, FromNotation("h1"))

	if err := validator.ValidateCastling(move.CastlingShort, b); err == nil {
		t.Errorf("the white shouldn't be able castle, got error: %v", err)
	}
}
