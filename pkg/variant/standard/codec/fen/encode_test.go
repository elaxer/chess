package fen

import (
	"testing"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/variant/standard/board"
	"github.com/elaxer/chess/pkg/variant/standardtest"
)

func TestEncode(t *testing.T) {
	factory := board.NewFactory()
	type args struct {
		board chess.Board
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"empty",
			args{factory.CreateEmpty(chess.SideWhite)},
			"8/8/8/8/8/8/8/8 w - - 0 1",
		},
		{
			"initial_position",
			args{factory.CreateFilled()},
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		},
		{
			"e4",
			args{standardtest.NewFromMoves([]string{"e4"})},
			"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1",
		},
		{
			"e4_d5_Nf3_Kd7",
			args{standardtest.NewFromMoves([]string{"e4", "d5", "Nf3", "Kd7"})},
			"rnbq1bnr/pppkpppp/8/3p4/4P3/5N2/PPPP1PPP/RNBQKB1R w KQ - 2 3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args.board); got != tt.want {
				t.Errorf("EncodeFEN() = \n%v want\n%v", got, tt.want)
			}
		})
	}
}
