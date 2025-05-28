package standard

import (
	"fmt"
	"strconv"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/metric"
	standardmetric "github.com/elaxer/chess/pkg/variant/standard/metric"
)

// EncodeFEN encodes the board position in FEN format.
// The FEN format is a standard notation for describing chess positions.
// It consists of six fields separated by spaces:
// 1. Piece placement
// 2. Active color
// 3. Castling availability
// 4. En passant target square
// 5. Halfmove clock
// 6. Fullmove number
func EncodeFEN(board chess.Board) string {
	return fmt.Sprintf(
		"%s %s %v %v %v %d",
		fenPosition(board.Squares()),
		board.Turn(),
		callMetric(standardmetric.CastlingsString, board),
		callMetric(standardmetric.EnPassantPosition, board),
		callMetric(standardmetric.FiftyMovesClock, board),
		len(board.MovesHistory())/2+1,
	)
}

func fenPosition(squares *chess.Squares) string {
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

			rowStr += fmt.Sprintf("%s", piece)
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
