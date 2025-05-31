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

func TestCheck(t *testing.T) {
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
			"check",
			args{
				standardtest.NewEmpty(chess.SideWhite, []standardtest.Placement{
					{Piece: piece.NewKing(chess.SideWhite), Position: position.FromString("a1")},
					{Piece: piece.NewKing(chess.SideBlack), Position: position.FromString("h8")},
					{Piece: piece.NewRook(chess.SideBlack), Position: position.FromString("a8")},
				}),
				chess.SideWhite,
			},
			state.Check,
		},
		{
			"check_bishop",
			args{
				standardtest.NewEmpty(chess.SideBlack, []standardtest.Placement{
					{Piece: piece.NewKing(chess.SideBlack), Position: position.FromString("e1")},
					{Piece: piece.NewKing(chess.SideWhite), Position: position.FromString("h8")},
					{Piece: piece.NewBishop(chess.SideWhite), Position: position.FromString("b4")},
				}),
				chess.SideBlack,
			},
			state.Check,
		},
		{
			"check_pawns",
			args{
				standardtest.NewEmpty(chess.SideBlack, []standardtest.Placement{
					{Piece: piece.NewKing(chess.SideBlack), Position: position.FromString("d4")},
					{Piece: piece.NewKing(chess.SideWhite), Position: position.FromString("c3")},
					{Piece: piece.NewBishop(chess.SideWhite), Position: position.FromString("e3")},
				}),
				chess.SideBlack,
			},
			state.Check,
		},
		{
			"no_check",
			args{
				standardtest.NewEmpty(chess.SideWhite, []standardtest.Placement{
					{Piece: piece.NewKing(chess.SideWhite), Position: position.FromString("d4")},
					{Piece: piece.NewKing(chess.SideBlack), Position: position.FromString("h8")},
					{Piece: piece.NewRook(chess.SideBlack), Position: position.FromString("a1")},
				}),
				chess.SideWhite,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rule.Check(tt.args.board, tt.args.side); got != tt.want {
				t.Errorf("Check() = %v, want %v", got, tt.want)
			}
		})
	}
}
