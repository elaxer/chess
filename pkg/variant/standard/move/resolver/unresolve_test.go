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
		move  *move.Normal
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
				&move.Normal{From: position.FromString("d1"), To: position.FromString("d4")},
				standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: piece.NewQueen(SideWhite), Position: position.FromString("d1")},
					{Piece: piece.NewQueen(SideWhite), Position: position.FromString("d8")},
				}),
			},
			position.Position{Rank: 1},
			false,
		},
		{
			"same_rank",
			args{
				&move.Normal{From: position.FromString("a1"), To: position.FromString("d1")},
				standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: piece.NewRook(SideBlack), Position: position.FromString("a1")},
					{Piece: piece.NewRook(SideBlack), Position: position.FromString("g1")},
				}),
			},
			position.Position{File: 1},
			false,
		},
		{
			"same_file_and_rank",
			args{
				&move.Normal{From: position.FromString("b7"), To: position.FromString("d5")},
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
				&move.Normal{From: position.FromString("g1"), To: position.FromString("e2")},
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
				&move.Normal{From: position.FromString("e2"), To: position.FromString("e4")},
				standardtest.NewEmpty(SideBlack, []standardtest.Placement{
					{Piece: piece.NewPawn(SideBlack), Position: position.FromString("e2")},
					{Piece: piece.NewPawn(SideBlack), Position: position.FromString("f2")},
				}),
			},
			position.NewNull(),
			false,
		},
		{
			"single_pawn_capture",
			args{
				&move.Normal{From: position.FromString("e2"), To: position.FromString("d1"), PieceNotation: piece.NotationPawn, IsCapture: true},
				standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: piece.NewPawn(SideWhite), Position: position.FromString("e2")},
					{Piece: piece.NewPawn(SideBlack), Position: position.FromString("d1")},
				}),
			},
			position.FromString("e"),
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
