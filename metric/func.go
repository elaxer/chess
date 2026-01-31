package metric

import "github.com/elaxer/chess"

// AllFuncs contains the default set of MetricFunc used to enrich visualizations/encoders.
var AllFuncs = []MetricFunc{
	HalfmoveCounter,
	FullmoveCounter,
	LastMove,
	Material,
	MaterialDifference,
}

// HalfmoveCounter returns a Metric with the number of half-moves made in the game.
func HalfmoveCounter(board chess.Board) Metric {
	return New("Halfmoves", len(board.MoveHistory()))
}

// FullmoveCounter returns a Metric with the number of full moves made in the game.
func FullmoveCounter(board chess.Board) Metric {
	return New("Full moves", (len(board.MoveHistory())+1)/2)
}

// LastMove returns a Metric with the last move made, or nil if no moves exist.
func LastMove(board chess.Board) Metric {
	movesCount := len(board.MoveHistory())
	if movesCount == 0 {
		return New("Last move", nil)
	}

	return New("Last move", board.MoveHistory()[movesCount-1].String())
}

// Material returns a Metric with material values for White and Black as a slice [white, black].
func Material(board chess.Board) Metric {
	callback := func(color chess.Color) uint16 {
		var weight uint16
		for piece := range board.Squares().GetAllPieces(color) {
			weight += (piece.Weight())
		}

		return weight
	}

	return New("Material value", []uint16{callback(chess.ColorWhite), callback(chess.ColorBlack)})
}

// MaterialDifference returns a Metric representing the material advantage (white - black).
func MaterialDifference(board chess.Board) Metric {
	material := Material(board).Value().([]uint16)

	return New("Material diff", int(material[0])-int(material[1]))
}
