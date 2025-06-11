package fen

import (
	"fmt"
	"iter"
	"strconv"
	"strings"

	"github.com/elaxer/chess"
	"github.com/elaxer/chess/metric"
	"github.com/elaxer/chess/position"
)

// Encoder encodes a chess board into a FEN string.
// It can also include additional metrics if provided.
type Encoder struct {
	MetricFuncs []metric.MetricFunc
}

// Encode encodes the given chess board into a FEN string.
// If the board is nil, it returns an empty string.
// If MetricFuncs are provided, it appends their results to the FEN string.
// The format of the FEN string is:
// <piece placement> <turn> [<metric1> <metric2> ...].
// If no metrics are provided, it will only include the piece placement and turn.
// If metric functions return nil, it will append a dash ("-") for that metric.
func (e *Encoder) Encode(board chess.Board) string {
	if board == nil {
		return ""
	}

	var fen strings.Builder
	fmt.Fprintf(&fen, "%s %s", EncodePiecePlacement(board.Squares()), board.Turn())

	if e.MetricFuncs != nil {
		for _, metricFunc := range e.MetricFuncs {
			fmt.Fprintf(&fen, " %v", callMetricFunc(metricFunc, board))
		}
	}

	return fen.String()
}

// EncodePiecePlacement encodes the piece placement of the given squares into a FEN string.
// It iterates through the squares by rows and encodes each row.
// Each row is represented by a string of piece string representation, with empty squares represented by numbers.
func EncodePiecePlacement(squares *chess.Squares) string {
	fen := ""
	for _, row := range squares.IterByRows(true) {
		fen += encodeRow(row) + "/"
	}

	return fen[:len(fen)-1]
}

func encodeRow(row iter.Seq2[position.File, chess.Piece]) string {
	var rowStr strings.Builder
	emptySquares := 0
	for _, piece := range row {
		if piece == nil {
			emptySquares++

			continue
		}

		if emptySquares > 0 {
			rowStr.WriteString(strconv.Itoa(emptySquares))
			emptySquares = 0
		}

		rowStr.WriteString(piece.String())
	}

	if emptySquares > 0 {
		rowStr.WriteString(strconv.Itoa(emptySquares))
	}

	return rowStr.String()
}

func callMetricFunc(metricFunc metric.MetricFunc, board chess.Board) any {
	metric := metricFunc(board)
	if metric == nil {
		return "-"
	}

	return metric.Value()
}
