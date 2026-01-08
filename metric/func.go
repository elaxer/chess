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
	return New("Full moves", len(board.MoveHistory())/2+1)
}

// LastMove returns a Metric with the last move made, or nil if no moves exist.
func LastMove(board chess.Board) Metric {
	movesCount := len(board.MoveHistory())
	if movesCount == 0 {
		return nil
	}

	return New("Last move", board.MoveHistory()[movesCount-1])
}

// Material returns a Metric with material values for White and Black as a slice [white, black].
func Material(board chess.Board) Metric {
	callback := func(color chess.Color) int {
		var weight int
		for piece := range board.Squares().GetAllPieces(color) {
			weight += int(piece.Weight())
		}

		return weight
	}

	return New("Material value", []int{callback(chess.ColorWhite), callback(chess.ColorBlack)})
}

// MaterialDifference returns a Metric representing the material advantage (white - black).
func MaterialDifference(board chess.Board) Metric {
	material := Material(board).Value().([]int)

	return New("Material advantage difference", material[0]-material[1])
}
