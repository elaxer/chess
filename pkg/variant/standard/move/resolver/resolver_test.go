package resolver_test

import (
	"testing"

	. "github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/move/resolver"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
	"github.com/elaxer/chess/pkg/variant/standarttest"
)

func TestResolveFrom(t *testing.T) {
	type args struct {
		from          position.Position
		to            position.Position
		pieceNotation PieceNotation
		board         Board
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
				from:          position.Position{},
				to:            position.FromNotation("e4"),
				pieceNotation: NotationPawn,
				board: standarttest.NewEmptyBoard(SideWhite, []standarttest.Placement{
					{Piece: piece.NewPawn(SideWhite), Position: position.FromNotation("d2")},
					{Piece: piece.NewPawn(SideWhite), Position: position.FromNotation("e2")},
					{Piece: piece.NewPawn(SideWhite), Position: position.FromNotation("f2")},
				}),
			},
			position.FromNotation("e2"),
			false,
		},
		{
			"same_file",
			args{
				from:          position.Position{File: position.FileA},
				to:            position.FromNotation("b8"),
				pieceNotation: NotationRook,
				board: standarttest.NewEmptyBoard(SideBlack, []standarttest.Placement{
					{Piece: piece.NewRook(SideBlack), Position: position.FromNotation("f8")},
					{Piece: piece.NewRook(SideBlack), Position: position.FromNotation("a8")},
				}),
			},
			position.FromNotation("a8"),
			false,
		},
		{
			"knights",
			args{
				from:          position.Position{File: position.FileG},
				to:            position.FromNotation("e2"),
				pieceNotation: NotationKnight,
				board: standarttest.NewEmptyBoard(SideWhite, []standarttest.Placement{
					{Piece: piece.NewKnight(SideWhite), Position: position.FromNotation("g1")},
					{Piece: piece.NewKnight(SideWhite), Position: position.FromNotation("c3")},
				}),
			},
			position.FromNotation("g1"),
			false,
		},
		{
			"same_rank",
			args{
				from:          position.Position{Rank: 1},
				to:            position.FromNotation("a5"),
				pieceNotation: NotationRook,
				board: standarttest.NewEmptyBoard(SideWhite, []standarttest.Placement{
					{Piece: piece.NewRook(SideWhite), Position: position.FromNotation("a1")},
					{Piece: piece.NewRook(SideWhite), Position: position.FromNotation("a8")},
				}),
			},
			position.FromNotation("a1"),
			false,
		},
		{
			"full_from",
			args{
				from:          position.FromNotation("f2"),
				to:            position.FromNotation("d4"),
				pieceNotation: NotationBishop,
				board: standarttest.NewEmptyBoard(SideBlack, []standarttest.Placement{
					{Piece: piece.NewBishop(SideBlack), Position: position.FromNotation("b2")},
					{Piece: piece.NewBishop(SideBlack), Position: position.FromNotation("f2")},
					{Piece: piece.NewBishop(SideBlack), Position: position.FromNotation("b6")},
				}),
			},
			position.FromNotation("f2"),
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := resolver.ResolveFrom(tt.args.from, tt.args.to, tt.args.pieceNotation, tt.args.board)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResolveFrom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ResolveFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnresolveFrom(t *testing.T) {
	type args struct {
		from  position.Position
		to    position.Position
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
				position.FromNotation("d1"),
				position.FromNotation("d4"),
				standarttest.NewEmptyBoard(SideWhite, []standarttest.Placement{
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
				position.FromNotation("a1"),
				position.FromNotation("d1"),
				standarttest.NewEmptyBoard(SideWhite, []standarttest.Placement{
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
				position.FromNotation("b7"),
				position.FromNotation("d5"),
				standarttest.NewEmptyBoard(SideWhite, []standarttest.Placement{
					{Piece: piece.NewBishop(SideWhite), Position: position.FromNotation("b7")},
					{Piece: piece.NewBishop(SideWhite), Position: position.FromNotation("f7")},
					{Piece: piece.NewBishop(SideWhite), Position: position.FromNotation("b3")},
				}),
			},
			position.FromNotation("b7"),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := resolver.UnresolveFrom(tt.args.from, tt.args.to, tt.args.board)
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
