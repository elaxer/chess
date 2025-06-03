package resolver_test

import (
	"testing"

	. "github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/move/move"
	"github.com/elaxer/chess/pkg/variant/standard/move/resolver"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
	"github.com/elaxer/chess/pkg/variant/standardtest"
)

func TestUnresolveFrom(t *testing.T) {
	type args struct {
		move  move.Piece
		board Board
	}
	tests := []struct {
		name    string
		args    args
		want    position.Position
		wantErr bool
	}{
		{
			"same_file",
			args{
				move.NewPiece(position.FromString("d1"), position.FromString("d4")),
				standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: piece.NewQueen(SideWhite), Position: position.FromString("d1")},
					{Piece: piece.NewQueen(SideWhite), Position: position.FromString("d8")},
				}),
			},
			position.Position{Rank: position.Rank1},
			false,
		},
		{
			"same_rank",
			args{
				move.NewPiece(position.FromString("a1"), position.FromString("d1")),
				standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: piece.NewRook(SideBlack), Position: position.FromString("a1")},
					{Piece: piece.NewRook(SideBlack), Position: position.FromString("g1")},
				}),
			},
			position.Position{File: position.FileA},
			false,
		},
		{
			"same_file_and_rank",
			args{
				move.NewPiece(position.FromString("b7"), position.FromString("d5")),
				standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: piece.NewBishop(SideWhite), Position: position.FromString("b7")},
					{Piece: piece.NewBishop(SideWhite), Position: position.FromString("f7")},
					{Piece: piece.NewBishop(SideWhite), Position: position.FromString("b3")},
				}),
			},
			position.FromString("b7"),
			false,
		},
		{
			"no_same_file_and_rank",
			args{
				move.NewPiece(position.FromString("g1"), position.FromString("e2")),
				standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: piece.NewKnight(SideWhite), Position: position.FromString("c3")},
					{Piece: piece.NewKnight(SideWhite), Position: position.FromString("g1")},
				}),
			},
			position.FromString("g"),
			false,
		},
		{
			"no_same_moves",
			args{
				move.NewPiece(position.FromString("e2"), position.FromString("e4")),
				standardtest.NewEmpty(SideBlack, []standardtest.Placement{
					{Piece: piece.NewPawn(SideBlack), Position: position.FromString("e2")},
					{Piece: piece.NewPawn(SideBlack), Position: position.FromString("f2")},
				}),
			},
			position.NewEmpty(),
			false,
		},
		{
			"single_pawn_capture",
			args{
				move.NewPiece(position.FromString("b7"), position.FromString("c8")),
				standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: piece.NewPawn(SideWhite), Position: position.FromString("b7")},
					{Piece: piece.NewPawn(SideBlack), Position: position.FromString("c8")},
				}),
			},
			position.NewEmpty(),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := resolver.UnresolveFrom(tt.args.move, tt.args.board)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnresolveFrom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UnresolveFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
