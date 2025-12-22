package visualizer

import (
	"fmt"
	"io"

	"github.com/elaxer/chess"
)

type Visualizer struct {
	Options Options
}

func (v *Visualizer) Visualize(board chess.Board, writer io.Writer) {
	backward := (v.Options.Orientation == OptionOrientationDefault) ||
		(v.Options.Orientation == OptionOrientationByTurn && board.Turn() == chess.SideWhite)

	for rank, row := range board.Squares().IterByRows(backward) {
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
