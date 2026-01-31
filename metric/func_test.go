package metric_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/chess/chesstest"
	"github.com/elaxer/chess/metric"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHalfmoveCounter(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		board     chess.Board
		wantValue any
	}{
		{
			"zero",
			&chesstest.BoardMock{MovesHistoryValue: []chess.Move{}},
			0,
		},
		{
			"one",
			&chesstest.BoardMock{MovesHistoryValue: []chess.Move{nil}},
			1,
		},
		{
			"odd",
			&chesstest.BoardMock{MovesHistoryValue: []chess.Move{nil, nil, nil, nil, nil}},
			5,
		},
		{
			"even",
			&chesstest.BoardMock{MovesHistoryValue: []chess.Move{nil, nil, nil, nil, nil, nil}},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := metric.HalfmoveCounter(tt.board)
			assert.Equal(t, tt.wantValue, got.Value())
		})
	}
}

func TestFullmoveCounter(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		board     chess.Board
		wantValue any
	}{
		{
			"zero",
			&chesstest.BoardMock{MovesHistoryValue: []chess.Move{}},
			0,
		},
		{
			"one",
			&chesstest.BoardMock{MovesHistoryValue: []chess.Move{nil}},
			1,
		},
		{
			"odd",
			&chesstest.BoardMock{MovesHistoryValue: []chess.Move{nil, nil, nil, nil, nil}},
			3,
		},
		{
			"even",
			&chesstest.BoardMock{MovesHistoryValue: []chess.Move{nil, nil, nil, nil, nil, nil}},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := metric.FullmoveCounter(tt.board)
			assert.Equal(t, tt.wantValue, got.Value())
		})
	}
}

func TestLastMove(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		board   chess.Board
		want    chess.Move
		wantNil bool
	}{
		{
			"zero",
			&chesstest.BoardMock{MovesHistoryValue: []chess.Move{}},
			nil,
			true,
		},
		{
			"A",
			&chesstest.BoardMock{MovesHistoryValue: []chess.Move{&chesstest.MoveResultMock{StringValue: "A"}}},
			&chesstest.MoveResultMock{StringValue: "A"},
			false,
		},
		{
			"C",
			&chesstest.BoardMock{
				MovesHistoryValue: []chess.Move{
					&chesstest.MoveResultMock{StringValue: "A"},
					&chesstest.MoveResultMock{StringValue: "B"},
					&chesstest.MoveResultMock{StringValue: "C"},
				},
			},
			&chesstest.MoveResultMock{StringValue: "C"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVal := metric.LastMove(tt.board).Value()

			if !tt.wantNil {
				require.NotNil(t, gotVal)
			} else {
				require.Nil(t, gotVal)

				return
			}

			assert.Equal(t, tt.want.String(), gotVal.(string))
		})
	}
}

func TestMaterial(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		board     chess.Board
		wantValue []uint16
	}{
		{
			"empty_board",
			&chesstest.BoardMock{
				SquaresValue: chesstest.MustSquaresFromPlacement(
					chess.PositionFromString("h8"),
					nil,
				),
			},
			[]uint16{0, 0},
		},
		{
			"only_one_side",
			&chesstest.BoardMock{
				SquaresValue: chesstest.MustSquaresFromPlacement(
					chess.PositionFromString("h8"),
					map[chess.Position]chess.Piece{
						chess.PositionFromString("a2"): &chesstest.PieceMock{ColorValue: chess.ColorWhite, WeightValue: 1},
						chess.PositionFromString("b2"): &chesstest.PieceMock{ColorValue: chess.ColorWhite, WeightValue: 2},
						chess.PositionFromString("c2"): &chesstest.PieceMock{ColorValue: chess.ColorWhite, WeightValue: 3},
					},
				),
			},
			[]uint16{6, 0},
		},
		{
			"both_sides",
			&chesstest.BoardMock{
				SquaresValue: chesstest.MustSquaresFromPlacement(
					chess.PositionFromString("h8"),
					map[chess.Position]chess.Piece{
						chess.PositionFromString("a2"): &chesstest.PieceMock{ColorValue: chess.ColorWhite, WeightValue: 1},
						chess.PositionFromString("b2"): &chesstest.PieceMock{ColorValue: chess.ColorWhite, WeightValue: 1},
						chess.PositionFromString("c2"): &chesstest.PieceMock{ColorValue: chess.ColorWhite, WeightValue: 1},
						chess.PositionFromString("d2"): &chesstest.PieceMock{ColorValue: chess.ColorWhite, WeightValue: 1},
						chess.PositionFromString("a1"): &chesstest.PieceMock{ColorValue: chess.ColorWhite, WeightValue: 5},
						chess.PositionFromString("d1"): &chesstest.PieceMock{ColorValue: chess.ColorWhite, WeightValue: 9},

						chess.PositionFromString("a7"): &chesstest.PieceMock{ColorValue: chess.ColorBlack, WeightValue: 1},
						chess.PositionFromString("b7"): &chesstest.PieceMock{ColorValue: chess.ColorBlack, WeightValue: 1},
						chess.PositionFromString("c7"): &chesstest.PieceMock{ColorValue: chess.ColorBlack, WeightValue: 1},
						chess.PositionFromString("d8"): &chesstest.PieceMock{ColorValue: chess.ColorBlack, WeightValue: 1},
						chess.PositionFromString("a8"): &chesstest.PieceMock{ColorValue: chess.ColorBlack, WeightValue: 5},
						chess.PositionFromString("c4"): &chesstest.PieceMock{ColorValue: chess.ColorBlack, WeightValue: 3},
						chess.PositionFromString("g8"): &chesstest.PieceMock{ColorValue: chess.ColorBlack, WeightValue: 9},
					},
				),
			},
			[]uint16{18, 21},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := metric.Material(tt.board).Value().([]uint16)
			assert.ElementsMatch(t, tt.wantValue, got)
		})
	}
}

func TestMaterialDifference(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		board     chess.Board
		wantValue int
	}{
		{
			"empty_board",
			&chesstest.BoardMock{
				SquaresValue: chesstest.MustSquaresFromPlacement(
					chess.PositionFromString("h8"),
					nil,
				),
			},
			0,
		},
		{
			"only_one_side",
			&chesstest.BoardMock{
				SquaresValue: chesstest.MustSquaresFromPlacement(
					chess.PositionFromString("h8"),
					map[chess.Position]chess.Piece{
						chess.PositionFromString("a2"): &chesstest.PieceMock{ColorValue: chess.ColorWhite, WeightValue: 1},
						chess.PositionFromString("b2"): &chesstest.PieceMock{ColorValue: chess.ColorWhite, WeightValue: 2},
						chess.PositionFromString("c2"): &chesstest.PieceMock{ColorValue: chess.ColorWhite, WeightValue: 3},
					},
				),
			},
			6,
		},
		{
			"white_superior",
			&chesstest.BoardMock{
				SquaresValue: chesstest.MustSquaresFromPlacement(
					chess.PositionFromString("g7"),
					map[chess.Position]chess.Piece{
						chess.PositionFromString("c7"): &chesstest.PieceMock{ColorValue: chess.ColorWhite, WeightValue: 7},
						chess.PositionFromString("d7"): &chesstest.PieceMock{ColorValue: chess.ColorWhite, WeightValue: 1},
						chess.PositionFromString("a7"): &chesstest.PieceMock{ColorValue: chess.ColorWhite, WeightValue: 5},
						chess.PositionFromString("c4"): &chesstest.PieceMock{ColorValue: chess.ColorWhite, WeightValue: 3},
						chess.PositionFromString("g7"): &chesstest.PieceMock{ColorValue: chess.ColorWhite, WeightValue: 9},

						chess.PositionFromString("d2"): &chesstest.PieceMock{ColorValue: chess.ColorBlack, WeightValue: 10},
						chess.PositionFromString("a1"): &chesstest.PieceMock{ColorValue: chess.ColorBlack, WeightValue: 5},
						chess.PositionFromString("d1"): &chesstest.PieceMock{ColorValue: chess.ColorBlack, WeightValue: 9},
					},
				),
			},
			1,
		},
		{
			"black_superior",
			&chesstest.BoardMock{
				SquaresValue: chesstest.MustSquaresFromPlacement(
					chess.PositionFromString("p16"),
					map[chess.Position]chess.Piece{
						chess.PositionFromString("a1"):  &chesstest.PieceMock{ColorValue: chess.ColorWhite, WeightValue: 5},
						chess.PositionFromString("d12"): &chesstest.PieceMock{ColorValue: chess.ColorWhite, WeightValue: 9},

						chess.PositionFromString("a7"): &chesstest.PieceMock{ColorValue: chess.ColorBlack, WeightValue: 1},
						chess.PositionFromString("b7"): &chesstest.PieceMock{ColorValue: chess.ColorBlack, WeightValue: 1},
						chess.PositionFromString("o9"): &chesstest.PieceMock{ColorValue: chess.ColorBlack, WeightValue: 100},
						chess.PositionFromString("d8"): &chesstest.PieceMock{ColorValue: chess.ColorBlack, WeightValue: 1},
						chess.PositionFromString("c4"): &chesstest.PieceMock{ColorValue: chess.ColorBlack, WeightValue: 3},
					},
				),
			},
			-92,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := metric.MaterialDifference(tt.board).Value().(int)
			require.Equal(t, tt.wantValue, got)
		})
	}
}
