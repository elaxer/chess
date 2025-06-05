package fen

import (
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/encoding/fen"
	"github.com/elaxer/chess/pkg/chess/metric"
	standardmetric "github.com/elaxer/chess/pkg/variant/standard/metric"
	"github.com/elaxer/chess/pkg/variant/standard/move/move"
)

// NewEncoder creates a new FEN encoder for the standard chess variant.
func NewEncoder() *fen.Encoder {
	return &fen.Encoder{
		MetricFuncs: []metric.MetricFunc{
			turn,
			castlingMetric,
			standardmetric.EnPassantTargetSquare,
			standardmetric.HalfmoveClock,
			metric.FullmoveCounter,
		},
	}
}

func turn(board chess.Board) metric.Metric {
	return metric.New("Turn", board.Turn())
}

func castlingMetric(board chess.Board) metric.Metric {
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
		return nil
	}

	return metric.New("Castling Ability", str)
}
