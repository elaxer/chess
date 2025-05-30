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
				&move.Normal{From: position.FromNotation("d1"), To: position.FromNotation("d4")},
				standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: piece.NewQueen(SideWhite), Position: position.FromNotation("d1")},
					{Piece: piece.NewQueen(SideWhite), Position: position.FromNotation("d8")},
				}),
			},
			position.Position{Rank: 1},
			false,
		},
		{
			"same_rank",
			args{
				&move.Normal{From: position.FromNotation("a1"), To: position.FromNotation("d1")},
				standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: piece.NewRook(SideBlack), Position: position.FromNotation("a1")},
					{Piece: piece.NewRook(SideBlack), Position: position.FromNotation("g1")},
				}),
			},
			position.Position{File: 1},
			false,
		},
		{
			"same_file_and_rank",
			args{
				&move.Normal{From: position.FromNotation("b7"), To: position.FromNotation("d5")},
				standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: piece.NewBishop(SideWhite), Position: position.FromNotation("b7")},
					{Piece: piece.NewBishop(SideWhite), Position: position.FromNotation("f7")},
					{Piece: piece.NewBishop(SideWhite), Position: position.FromNotation("b3")},
				}),
			},
			position.FromNotation("b7"),
			false,
		},
		{
			"no_same_file_and_rank",
			args{
				&move.Normal{From: position.FromNotation("g1"), To: position.FromNotation("e2")},
				standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: piece.NewKnight(SideWhite), Position: position.FromNotation("c3")},
					{Piece: piece.NewKnight(SideWhite), Position: position.FromNotation("g1")},
				}),
			},
			position.FromNotation("g"),
			false,
		},
		{
			"no_same_moves",
			args{
				&move.Normal{From: position.FromNotation("e2"), To: position.FromNotation("e4")},
				standardtest.NewEmpty(SideBlack, []standardtest.Placement{
					{Piece: piece.NewPawn(SideBlack), Position: position.FromNotation("e2")},
					{Piece: piece.NewPawn(SideBlack), Position: position.FromNotation("f2")},
				}),
			},
			position.NewNull(),
			false,
		},
		{
			"single_pawn_capturing",
			args{
				&move.Normal{From: position.FromNotation("e2"), To: position.FromNotation("d1"), PieceNotation: piece.NotationPawn, IsCapture: true},
				standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: piece.NewPawn(SideWhite), Position: position.FromNotation("e2")},
					{Piece: piece.NewPawn(SideBlack), Position: position.FromNotation("d1")},
				}),
			},
			position.FromNotation("e"),
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
