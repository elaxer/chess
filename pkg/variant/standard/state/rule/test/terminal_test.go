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

// todo test opposite turn
func TestCheckmate(t *testing.T) {
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
			"checkmate",
			args{
				standardtest.NewEmpty(chess.SideWhite, []standardtest.Placement{
					{Piece: piece.NewKing(chess.SideWhite), Position: position.FromString("a1")},
					{Piece: piece.NewKing(chess.SideBlack), Position: position.FromString("h8")},
					{Piece: piece.NewRook(chess.SideBlack), Position: position.FromString("a8")},
					{Piece: piece.NewRook(chess.SideBlack), Position: position.FromString("b8")},
				}),
				chess.SideWhite,
			},
			state.Checkmate,
		},
		{
			// no checkmate because the black king can capture the threatening rook
			"no_checkmate",
			args{
				standardtest.NewEmpty(chess.SideWhite, []standardtest.Placement{
					{Piece: piece.NewKing(chess.SideWhite), Position: position.FromString("a1")},
					{Piece: piece.NewKing(chess.SideBlack), Position: position.FromString("h8")},
					{Piece: piece.NewRook(chess.SideBlack), Position: position.FromString("a2")},
					{Piece: piece.NewRook(chess.SideBlack), Position: position.FromString("b8")},
				}),
				chess.SideWhite,
			},

			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rule.Checkmate(tt.args.board, tt.args.side); got != tt.want {
				t.Errorf("Checkmate() = %v, want %v", got, tt.want)
			}
		})
	}
}
