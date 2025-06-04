package rule_test

import (
	"testing"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
	"github.com/elaxer/chess/pkg/variant/standard/standardtest"
	"github.com/elaxer/chess/pkg/variant/standard/state/rule"
	"github.com/elaxer/chess/pkg/variant/standard/state/state"
)

func TestStalemate(t *testing.T) {
	type args struct {
		board chess.Board
		side  chess.Side
	}
	tests := []struct {
		name string
		args args
		want chess.State
	}{
		{
			"stalemate",
			args{
				standardtest.NewEmpty(chess.SideWhite, []standardtest.Placement{
					{Piece: piece.NewKing(chess.SideWhite), Position: position.FromString("a8")},
					{Piece: piece.NewKing(chess.SideBlack), Position: position.FromString("b6")},
					{Piece: piece.NewQueen(chess.SideBlack), Position: position.FromString("c7")},
				}),
				chess.SideWhite,
			},
			state.Stalemate,
		},
		{
			"no_stalemate",
			args{
				standardtest.NewEmpty(chess.SideBlack, []standardtest.Placement{
					{Piece: piece.NewKing(chess.SideWhite), Position: position.FromString("a8")},
					{Piece: piece.NewQueen(chess.SideWhite), Position: position.FromString("c7")},
					{Piece: piece.NewKing(chess.SideBlack), Position: position.FromString("b6")},
				}),
				chess.SideBlack,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rule.Stalemate(tt.args.board, tt.args.side); got != tt.want {
				t.Errorf("Stalemate() = %v, want %v", got, tt.want)
			}
		})
	}
}
