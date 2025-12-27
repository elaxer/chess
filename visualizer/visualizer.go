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
		(v.Options.Orientation == OptionOrientationByTurn && board.Turn() == chess.SideWhite)

	for rank, row := range board.Squares().IterOverRows(backward) {
		if v.Options.DisplayPositions {
			//nolint:errcheck
			fmt.Fprintf(writer, "%d ", rank)
		}
		for _, piece := range row {
			if piece == nil {
				//nolint:errcheck
				fmt.Fprint(writer, ". ")
			} else {
				//nolint:errcheck
				fmt.Fprintf(writer, "%s ", piece)
			}
		}

		//nolint:errcheck
		fmt.Fprintln(writer)
	}

	if v.Options.DisplayPositions {
		v.displayFilePositions(board, writer)
	}

	v.displayMetrics(board, writer)
}

// Fprintln is like Fprint but adds a newline after the board representation.
func (v *Visualizer) Fprintln(writer io.Writer, board chess.Board) {
	v.Fprint(writer, board)
	//nolint:errcheck
	fmt.Fprintln(writer)
}

func (v *Visualizer) displayFilePositions(board chess.Board, writer io.Writer) {
	//nolint:errcheck
	fmt.Fprintf(writer, "  ")
	for file := range board.Squares().EdgePosition().File {
		//nolint:errcheck
		fmt.Fprintf(writer, "%s ", file+1)
	}
}

func (v *Visualizer) displayMetrics(board chess.Board, writer io.Writer) {
	for _, metricFunc := range v.Options.MetricFuncs {
		if metric := metricFunc(board); metric != nil {
			//nolint:errcheck
			fmt.Fprintf(writer, "\n%s", metric)
		}
	}
}
