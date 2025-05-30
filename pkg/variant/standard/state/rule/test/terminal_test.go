package rule_test

import (
	"testing"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
	"github.com/elaxer/chess/pkg/variant/standard/state/rule"
	"github.com/elaxer/chess/pkg/variant/standard/state/state"
	"github.com/elaxer/chess/pkg/variant/standardtest"
)

// todo test opposite turn
func TestMate(t *testing.T) {
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
			"mate",
			args{
				standardtest.NewEmpty(chess.SideWhite, []standardtest.Placement{
					{Piece: piece.NewKing(chess.SideWhite), Position: position.FromNotation("a1")},
					{Piece: piece.NewKing(chess.SideBlack), Position: position.FromNotation("h8")},
					{Piece: piece.NewRook(chess.SideBlack), Position: position.FromNotation("a8")},
					{Piece: piece.NewRook(chess.SideBlack), Position: position.FromNotation("b8")},
				}),
				chess.SideWhite,
			},
			state.Mate,
		},
		{
			// no mate because the black king can capture the threatening rook
			"no_mate",
			args{
				standardtest.NewEmpty(chess.SideWhite, []standardtest.Placement{
					{Piece: piece.NewKing(chess.SideWhite), Position: position.FromNotation("a1")},
					{Piece: piece.NewKing(chess.SideBlack), Position: position.FromNotation("h8")},
					{Piece: piece.NewRook(chess.SideBlack), Position: position.FromNotation("a2")},
					{Piece: piece.NewRook(chess.SideBlack), Position: position.FromNotation("b8")},
				}),
				chess.SideWhite,
			},

			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rule.Mate(tt.args.board, tt.args.side); got != tt.want {
				t.Errorf("Mate() = %v, want %v", got, tt.want)
			}
		})
	}
}
