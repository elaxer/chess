// Package visualizer renders a human-readable textual representation of a chess.Board.
// It supports orientation options, optional display of rank/file indices, and the inclusion
// of metric functions to show additional information below the board. Visualizer is
// intended for debugging, tests, or simple CLI display.
package visualizer

import (
	"fmt"
	"io"

	"github.com/elaxer/chess"
)

// Visualizer renders a chess board to an io.Writer according to the provided Options.
// It prints ranks and files as text and optionally shows metrics below the board.
type Visualizer struct {
	Options Options
}

// Fprint writes a textual representation of the board to the provided writer.
// The output format depends on the Visualizer Options (orientation, positions display and metrics).
func (v *Visualizer) Fprint(writer io.Writer, board chess.Board) {
	backward := (v.Options.Orientation == OptionOrientationDefault) ||
		(v.Options.Orientation == OptionOrientationByTurn && board.Turn() == chess.ColorWhite)

	for rank, row := range board.Squares().IterOverRows(backward) {
		if v.Options.DisplayPositions {
			//nolint:errcheck
			fmt.Fprintf(writer, "%d ", rank)
		}
		for file, piece := range row {
			if piece == nil {
				//nolint:errcheck
				fmt.Fprint(writer, ".")
			} else {
				//nolint:errcheck
				fmt.Fprintf(writer, "%s", piece)
			}

			if file != board.Squares().EdgePosition().File {
				//nolint:errcheck
				fmt.Fprint(writer, " ")
			}
		}

		if rank != v.edgeRank(board.Squares(), backward) {
			//nolint:errcheck
			fmt.Fprintln(writer)
		}
	}

	if v.Options.DisplayPositions {
		v.displayFilePositions(writer, board)
	}

	v.displayMetrics(writer, board)
}

// Fprintln is like Fprint but adds a newline after the board representation.
func (v *Visualizer) Fprintln(writer io.Writer, board chess.Board) {
	v.Fprint(writer, board)
	//nolint:errcheck
	fmt.Fprintln(writer)
}

func (v *Visualizer) displayFilePositions(writer io.Writer, board chess.Board) {
	//nolint:errcheck
	fmt.Fprint(writer, "\n  ")
	for file := range board.Squares().EdgePosition().File {
		f := file + 1
		//nolint:errcheck
		fmt.Fprint(writer, f.String())
		if f != board.Squares().EdgePosition().File {
			//nolint:errcheck
			fmt.Fprint(writer, " ")
		}
	}
}

func (v *Visualizer) displayMetrics(writer io.Writer, board chess.Board) {
	for _, metricFunc := range v.Options.MetricFuncs {
		if metric := metricFunc(board); metric != nil {
			//nolint:errcheck
			fmt.Fprintf(writer, "\n%s", metric)
		}
	}
}

func (v *Visualizer) edgeRank(squares *chess.Squares, backward bool) chess.Rank {
	if backward {
		return chess.Rank1

	}

	return squares.EdgePosition().Rank
}
