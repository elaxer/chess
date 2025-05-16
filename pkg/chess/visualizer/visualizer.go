package visualizer

import (
	"fmt"
	"io"
	"slices"

	"github.com/elaxer/chess/pkg/chess"
)

type Visualizer struct {
	Options Options
}

func (v *Visualizer) Visualize(board chess.Board, writer io.Writer) {
	squares := slices.All(board.Squares().Items())
	reverse := (v.Options.Orientation == OptionOrientationDefault) || (v.Options.Orientation == OptionOrientationByTurn && board.Turn() == chess.SideWhite)
	if reverse {
		squares = slices.Backward(board.Squares().Items())
	}

	for i, row := range squares {
		if v.Options.Positions {
			fmt.Fprintf(writer, "%d ", i+1)
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
