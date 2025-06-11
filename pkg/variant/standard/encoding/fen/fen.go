package fen

import (
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/encoding/fen"
	"github.com/elaxer/chess/pkg/chess/metric"
	"github.com/elaxer/chess/pkg/variant/standard/board"
	standardmetric "github.com/elaxer/chess/pkg/variant/standard/metric"
	"github.com/elaxer/chess/pkg/variant/standard/move/move"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
)

// NewEncoder creates a new FEN encoder for the standard chess variant.
// Encoder.Encode() method will return a FEN string representing the current state of the board,
// including the turn, castling rights, en passant target square, halfmove clock, and fullmove counter.
//
// See github.com/elaxer/chess/pkg/chess/encoding/fen/encoder for more details.
func NewEncoder() *fen.Encoder {
	return &fen.Encoder{
		MetricFuncs: []metric.MetricFunc{
			castlingMetric,
			standardmetric.EnPassantTargetSquare,
			standardmetric.HalfmoveClock,
			metric.FullmoveCounter,
		},
	}
}

// NewDecoder creates a new FEN decoder for the standard chess variant.
// See github.com/elaxer/chess/pkg/chess/encoding/fen/decoder for more details.
func NewDecoder() *fen.Decoder {
	return fen.NewDecoder(board.NewFactory(), piece.NewFactory())
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
