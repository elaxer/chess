package resolver_test

import (
	"testing"

	. "github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/move"
	resolver "github.com/elaxer/chess/pkg/variant/standard/moveresolver"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
	"github.com/elaxer/chess/pkg/variant/standardtest"
)

func TestResolveNormal_From(t *testing.T) {
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
				move: &move.Normal{
					To:            position.FromNotation("e4"),
					PieceNotation: piece.NotationPawn,
				},
				board: standardtest.NewEmpty(SideWhite, []standardtest.Placement{
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
				move: &move.Normal{
					From:          position.Position{File: position.FileA},
					To:            position.FromNotation("b8"),
					PieceNotation: piece.NotationRook,
				},
				board: standardtest.NewEmpty(SideBlack, []standardtest.Placement{
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
				move: &move.Normal{
					From:          position.Position{File: position.FileG},
					To:            position.FromNotation("e2"),
					PieceNotation: piece.NotationKnight,
				},
				board: standardtest.NewEmpty(SideWhite, []standardtest.Placement{
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
				move: &move.Normal{
					From:          position.Position{Rank: 1},
					To:            position.FromNotation("a5"),
					PieceNotation: piece.NotationRook,
				},
				board: standardtest.NewEmpty(SideWhite, []standardtest.Placement{
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
				move: &move.Normal{
					From:          position.FromNotation("f2"),
					To:            position.FromNotation("d4"),
					PieceNotation: piece.NotationBishop,
				},
				board: standardtest.NewEmpty(SideBlack, []standardtest.Placement{
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
			got, err := resolver.ResolveNormal(tt.args.move, tt.args.board, tt.args.board.Turn())
			if (err != nil) != tt.wantErr {
				t.Errorf("ResolveFrom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.From != tt.want {
				t.Errorf("ResolveFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnresolveFrom(t *testing.T) {
	type args struct {
		from, to position.Position
		board    Board
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
				position.FromNotation("a1"),
				position.FromNotation("d1"),
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
				position.FromNotation("b7"),
				position.FromNotation("d5"),
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
			"not_same_file_and_rank",
			args{
				position.FromNotation("g1"),
				position.FromNotation("e2"),
				standardtest.NewEmpty(SideWhite, []standardtest.Placement{
					{Piece: piece.NewKnight(SideWhite), Position: position.FromNotation("c3")},
					{Piece: piece.NewKnight(SideWhite), Position: position.FromNotation("g1")},
				}),
			},
			position.FromNotation("g"),
			false,
		},
		{
			"not_same_moves",
			args{
				position.FromNotation("e2"),
				position.FromNotation("e4"),
				standardtest.NewEmpty(SideBlack, []standardtest.Placement{
					{Piece: piece.NewPawn(SideBlack), Position: position.FromNotation("e2")},
					{Piece: piece.NewPawn(SideBlack), Position: position.FromNotation("f2")},
				}),
			},
			position.NewNull(),
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
