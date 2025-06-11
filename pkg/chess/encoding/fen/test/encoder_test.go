package fen_test

import (
	"testing"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/chesstest"
	"github.com/elaxer/chess/pkg/chess/encoding/fen"
	"github.com/elaxer/chess/pkg/chess/metric"
	"github.com/elaxer/chess/pkg/chess/position"
)

func TestEncoder_Encode(t *testing.T) {
	encoder := new(fen.Encoder)

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
			args{chesstest.DecodeFEN8x8("8/8/8/8/8/8/8/8")},
			"8/8/8/8/8/8/8/8 w",
		},
		{
			"white",
			args{chesstest.DecodeFEN8x8("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w")},
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w",
		},
		{
			"black",
			args{chesstest.DecodeFEN8x8("3b2n1/pp2p1p1/2k2r2/3P4/8/2N1B1q1/3P4/R3K3 b")},
			"3b2n1/pp2p1p1/2k2r2/3P4/8/2N1B1q1/3P4/R3K3 b",
		},
		{
			"valid",
			args{chesstest.DecodeFEN8x8("3r2k1/pRp4p/2R3p1/8/3K4/P4r2/2P4P/1N6")},
			"3r2k1/pRp4p/2R3p1/8/3K4/P4r2/2P4P/1N6 w",
		},
		{
			"6x6",
			args{chesstest.DecodeFEN("rnq1kb/pp1ppp/6/6/PPP1PP/RNQK1B", position.FromString("f6"))},
			"rnq1kb/pp1ppp/6/6/PPP1PP/RNQK1B w",
		},
		{
			"12x8",
			args{chesstest.DecodeFEN("2r1k2bnq1p/3p4p1P1/2n3B1P3/1P1b2N3p1/2Q1P1p3R1/3K3P2N1/1p2B2P2p1/12", position.FromString("l8"))},
			"2r1k2bnq1p/3p4p1P1/2n3B1P3/1P1b2N3p1/2Q1P1p3R1/3K3P2N1/1p2B2P2p1/12 w",
		},
		{
			"7x13",
			args{chesstest.DecodeFEN(
				"1q2b2/2p1P1n/1r2B2/2N2pk/3r1PB/1Q2n1K/3P3/1P3r1/2N1k1P/2p1bP1/1R2K2/2q2p1/2B1n2",
				position.FromString("g13"),
			)},
			"1q2b2/2p1P1n/1r2B2/2N2pk/3r1PB/1Q2n1K/3P3/1P3r1/2N1k1P/2p1bP1/1R2K2/2q2p1/2B1n2 w",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encoder.Encode(tt.args.board); got != tt.want {
				t.Errorf("Encoder.Encode() = \n%v want\n%v", got, tt.want)
			}
		})
	}
}

func TestEncoder_EncodeWithMetricFuncs(t *testing.T) {
	board := chesstest.DecodeFEN8x8("3b2n1/pp2p1p1/2k2r2/3P4/8/2N1B1q1/3P4/R3K3 b")
	encoder := &fen.Encoder{MetricFuncs: []metric.MetricFunc{
		func(board chess.Board) metric.Metric {
			return metric.New("test_metric", 42)
		},
		func(board chess.Board) metric.Metric {
			return metric.New("test_metric_2", "val")
		},
		func(board chess.Board) metric.Metric {
			return nil
		},
	}}

	const expected = "3b2n1/pp2p1p1/2k2r2/3P4/8/2N1B1q1/3P4/R3K3 b 42 val -"
	if got := encoder.Encode(board); got != expected {
		t.Errorf("Encoder.Encode() = \n%v want\n%v", got, expected)
	}
}
