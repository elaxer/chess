package visualizer

import (
	"fmt"
	"io"

	"github.com/elaxer/chess/pkg/chess"
)

type Visualizer struct {
	Options Options
}

func (v *Visualizer) Visualize(board chess.Board, writer io.Writer) {
	backward := (v.Options.Orientation == OptionOrientationDefault) || (v.Options.Orientation == OptionOrientationByTurn && board.Turn() == chess.SideWhite)

	for rank, row := range board.Squares().IterByRows(backward) {
		if v.Options.Positions {
			fmt.Fprintf(writer, "%d ", rank)
		}
		for _, piece := range row {
			if piece == nil {
				fmt.Fprint(writer, ". ")
			} else {
				fmt.Fprintf(writer, "%s ", piece)
			}
		}

		fmt.Fprintln(writer)
	}

	if v.Options.Positions {
		fmt.Fprintf(writer, "  ")
		for file := range board.Squares().EdgePosition().File {
			fmt.Fprintf(writer, "%s ", file+1)
		}
	}

	for _, metricFunc := range v.Options.MetricFuncs {
		if metric := metricFunc(board); metric != nil {
			fmt.Fprintf(writer, "\n%s", metric)
		}
	}

	fmt.Fprintln(writer)
}
