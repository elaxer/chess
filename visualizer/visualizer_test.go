package visualizer_test

import (
	"strings"
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/chess/chesstest"
	"github.com/elaxer/chess/metric"
	"github.com/elaxer/chess/visualizer"
	"github.com/stretchr/testify/assert"
)

func TestVisualizer_Fprint(t *testing.T) {
	tests := []struct {
		name  string
		vis   *visualizer.Visualizer
		board chess.Board
		want  string
	}{
		{
			"empty_visualizer",
			&visualizer.Visualizer{},
			&chesstest.BoardMock{
				SquaresValue: chesstest.MustSquaresFromPlacement(chess.PositionFromString("h8"), map[chess.Position]chess.Piece{
					chess.PositionFromString("a1"): &chesstest.PieceMock{StringValue: "R"},
					chess.PositionFromString("h8"): &chesstest.PieceMock{StringValue: "r"},
					chess.PositionFromString("d4"): &chesstest.PieceMock{StringValue: "P"},
					chess.PositionFromString("d6"): &chesstest.PieceMock{StringValue: "p"},
				}),
			},
			`. . . . . . . r
. . . . . . . .
. . . p . . . .
. . . . . . . .
. . . P . . . .
. . . . . . . .
. . . . . . . .
R . . . . . . .`,
		},
		{
			"with_positions",
			&visualizer.Visualizer{
				Options: visualizer.Options{DisplayPositions: true},
			},
			&chesstest.BoardMock{
				SquaresValue: chesstest.MustSquaresFromPlacement(chess.PositionFromString("h8"), map[chess.Position]chess.Piece{
					chess.PositionFromString("a1"): &chesstest.PieceMock{StringValue: "R"},
					chess.PositionFromString("h8"): &chesstest.PieceMock{StringValue: "r"},
					chess.PositionFromString("d4"): &chesstest.PieceMock{StringValue: "P"},
					chess.PositionFromString("d6"): &chesstest.PieceMock{StringValue: "p"},
				}),
			},
			`8 . . . . . . . r
7 . . . . . . . .
6 . . . p . . . .
5 . . . . . . . .
4 . . . P . . . .
3 . . . . . . . .
2 . . . . . . . .
1 R . . . . . . .
  a b c d e f g h`,
		},
		{
			"orientation_reversed",
			&visualizer.Visualizer{
				Options: visualizer.Options{DisplayPositions: true, Orientation: visualizer.OptionOrientationReversed},
			},
			&chesstest.BoardMock{
				SquaresValue: chesstest.MustSquaresFromPlacement(chess.PositionFromString("h8"), map[chess.Position]chess.Piece{
					chess.PositionFromString("d6"): &chesstest.PieceMock{StringValue: "p"},
					chess.PositionFromString("a1"): &chesstest.PieceMock{StringValue: "R"},
					chess.PositionFromString("h8"): &chesstest.PieceMock{StringValue: "r"},
					chess.PositionFromString("d4"): &chesstest.PieceMock{StringValue: "P"},
				}),
				TurnValue: chess.ColorWhite,
			},
			`1 R . . . . . . .
2 . . . . . . . .
3 . . . . . . . .
4 . . . P . . . .
5 . . . . . . . .
6 . . . p . . . .
7 . . . . . . . .
8 . . . . . . . r
  a b c d e f g h`,
		},
		{
			"orientation_by_turn",
			&visualizer.Visualizer{
				Options: visualizer.Options{DisplayPositions: true, Orientation: visualizer.OptionOrientationByTurn},
			},
			&chesstest.BoardMock{
				SquaresValue: chesstest.MustSquaresFromPlacement(chess.PositionFromString("h8"), map[chess.Position]chess.Piece{
					chess.PositionFromString("d6"): &chesstest.PieceMock{StringValue: "p"},
					chess.PositionFromString("a1"): &chesstest.PieceMock{StringValue: "R"},
					chess.PositionFromString("h8"): &chesstest.PieceMock{StringValue: "r"},
					chess.PositionFromString("d4"): &chesstest.PieceMock{StringValue: "P"},
				}),
				TurnValue: chess.ColorBlack,
			},
			`1 R . . . . . . .
2 . . . . . . . .
3 . . . . . . . .
4 . . . P . . . .
5 . . . . . . . .
6 . . . p . . . .
7 . . . . . . . .
8 . . . . . . . r
  a b c d e f g h`,
		},
		{
			"with_metrics",
			&visualizer.Visualizer{
				Options: visualizer.Options{
					DisplayPositions: true,
					Orientation:      visualizer.OptionOrientationByTurn,
					MetricFuncs:      metric.AllFuncs,
				},
			},
			&chesstest.BoardMock{
				SquaresValue: chesstest.MustSquaresFromPlacement(chess.PositionFromString("h8"), map[chess.Position]chess.Piece{
					chess.PositionFromString("c1"): &chesstest.PieceMock{
						StringValue: "R",
						ColorValue:  chess.ColorWhite,
						WeightValue: 6,
					},
					chess.PositionFromString("d4"): &chesstest.PieceMock{
						StringValue: "K",
						ColorValue:  chess.ColorWhite,
						WeightValue: 7,
					},
					chess.PositionFromString("h6"): &chesstest.PieceMock{
						StringValue: "p",
						ColorValue:  chess.ColorBlack,
						WeightValue: 5,
					},
					chess.PositionFromString("h8"): &chesstest.PieceMock{
						StringValue: "r",
						ColorValue:  chess.ColorBlack,
						WeightValue: 6,
					},
				}),
				TurnValue: chess.ColorBlack,
				MovesHistoryValue: []chess.MoveResult{
					&chesstest.MoveResultMock{},
					&chesstest.MoveResultMock{},
					&chesstest.MoveResultMock{StringValue: "Kd4"},
				},
			},
			`1 . . R . . . . .
2 . . . . . . . .
3 . . . . . . . .
4 . . . K . . . .
5 . . . . . . . .
6 . . . . . . . p
7 . . . . . . . .
8 . . . . . . . r
  a b c d e f g h
Halfmoves: 3
Full moves: 2
Last move: Kd4
Material value: [13 11]
Material diff: 2`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var str strings.Builder

			tt.vis.Fprint(&str, tt.board)
			assert.Equal(t, tt.want, str.String())
		})
	}
}
