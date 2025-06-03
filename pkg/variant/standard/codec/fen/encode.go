package fen

import (
	"fmt"
	"strconv"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/metric"
	standardmetric "github.com/elaxer/chess/pkg/variant/standard/metric"
	"github.com/elaxer/chess/pkg/variant/standard/move/move"
)

// Encode encodes the board position in FEN format.
// The FEN format is a standard notation for describing chess positions.
// It consists of six fields separated by spaces:
// 1. Piece placement
// 2. Active color
// 3. Castling availability
// 4. En passant target square
// 5. Halfmove clock
// 6. Fullmove number
func Encode(board chess.Board) string {
	return fmt.Sprintf(
		"%s %s %v %v %v %d",
		EncodePiecePlacements(board.Squares()),
		board.Turn(),
		castlingAbility(board),
		callMetric(standardmetric.EnPassantTargetSquare, board),
		callMetric(standardmetric.HalfmoveClock, board),
		callMetric(metric.FullmoveCounter, board),
	)
}

func EncodePiecePlacements(squares *chess.Squares) string {
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

func castlingAbility(board chess.Board) string {
	theoretical := standardmetric.CastlingAbility(board).Value().(standardmetric.Castlings)["theoretical"]
	str := ""
	if theoretical[chess.SideWhite][move.CastlingShort] {
		str += "K"
	}
	if theoretical[chess.SideWhite][move.CastlingLong] {
		str += "Q"
	}
	if theoretical[chess.SideBlack][move.CastlingShort] {
		str += "k"
	}
	if theoretical[chess.SideBlack][move.CastlingLong] {
		str += "q"
	}

	if str == "" {
		return "-"
	}

	return str
}

func callMetric(metricFunc metric.MetricFunc, board chess.Board) any {
	metric := metricFunc(board)
	if metric == nil {
		return "-"
	}

	return metric.Value()
}
