package fen

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/metric"
)

type Encoder struct {
	MetricFuncs []metric.MetricFunc
}

func (e *Encoder) Encode(board chess.Board) string {
	var fen strings.Builder
	fmt.Fprint(&fen, encodePiecePlacement(board.Squares()))

	if e.MetricFuncs != nil {
		for _, metricFunc := range e.MetricFuncs {
			fmt.Fprintf(&fen, " %v", callMetric(metricFunc, board))
		}
	}

	return fen.String()
}

func encodePiecePlacement(squares *chess.Squares) string {
	fen := ""
	for _, row := range squares.IterByRows(true) {
		rowStr := ""
		emptySquares := 0
		for _, piece := range row {
			if piece == nil {
				emptySquares++

				continue
			}

			if emptySquares > 0 {
				rowStr += strconv.Itoa(emptySquares)
			}
			emptySquares = 0

			rowStr += piece.String()
		}

		if emptySquares > 0 {
			rowStr += strconv.Itoa(emptySquares)
		}

		fen += rowStr + "/"
	}

	return fen[:len(fen)-1]
}

func callMetric(metricFunc metric.MetricFunc, board chess.Board) any {
	metric := metricFunc(board)
	if metric == nil {
		return "-"
	}

	return metric.Value()
}
