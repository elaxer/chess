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

func TestResolveFrom(t *testing.T) {
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
			"empty_from",
			args{
				move: move.NewNormal(position.NewEmpty(), position.FromString("e4"), piece.NotationPawn),
				board: standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: piece.NewPawn(SideWhite), Position: position.FromString("d2")},
					{Piece: piece.NewPawn(SideWhite), Position: position.FromString("e2")},
					{Piece: piece.NewPawn(SideWhite), Position: position.FromString("f2")},
				}),
			},
			position.FromString("e2"),
			false,
		},
		{
			"same_file",
			args{
				move: move.NewNormal(position.FromString("a"), position.FromString("b8"), piece.NotationRook),
				board: standardtest.NewEmpty(SideBlack, []standardtest.Placement{
					{Piece: piece.NewRook(SideBlack), Position: position.FromString("f8")},
					{Piece: piece.NewRook(SideBlack), Position: position.FromString("a8")},
				}),
			},
			position.FromString("a8"),
			false,
		},
		{
			"knights",
			args{
				move: move.NewNormal(position.FromString("g"), position.FromString("e2"), piece.NotationKnight),
				board: standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: piece.NewKnight(SideWhite), Position: position.FromString("g1")},
					{Piece: piece.NewKnight(SideWhite), Position: position.FromString("c3")},
				}),
			},
			position.FromString("g1"),
			false,
		},
		{
			"same_rank",
			args{
				move: move.NewNormal(position.FromString("1"), position.FromString("a5"), piece.NotationRook),
				board: standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: piece.NewRook(SideWhite), Position: position.FromString("a1")},
					{Piece: piece.NewRook(SideWhite), Position: position.FromString("a8")},
				}),
			},
			position.FromString("a1"),
			false,
		},
		{
			"full_from",
			args{
				move: move.NewNormal(position.FromString("f2"), position.FromString("d4"), piece.NotationBishop),
				board: standardtest.NewEmpty(SideBlack, []standardtest.Placement{
					{Piece: piece.NewBishop(SideBlack), Position: position.FromString("b2")},
					{Piece: piece.NewBishop(SideBlack), Position: position.FromString("f2")},
					{Piece: piece.NewBishop(SideBlack), Position: position.FromString("b6")},
				}),
			},
			position.FromString("f2"),
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := resolver.ResolveFrom(tt.args.move.Piece, tt.args.move.PieceNotation, tt.args.board, tt.args.board.Turn())
			if (err != nil) != tt.wantErr {
				t.Errorf("ResolveNormal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ResolveNormal() = %v, want %v", got, tt.want)
			}
		})
	}
}
