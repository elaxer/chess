package metric

import "github.com/elaxer/chess"

var AllFuncs = []MetricFunc{
	HalfmoveCounter,
	FullmoveCounter,
	LastMove,
	Material,
	MaterialDifference,
}

func HalfmoveCounter(board chess.Board) Metric {
	return New("Halfmoves", len(board.MovesHistory()))
}

func FullmoveCounter(board chess.Board) Metric {
	moves := len(board.MovesHistory())
	fullmove := moves / 2

	if moves%2 != 0 || moves == 0 {
		fullmove++
	}

	return New("Full moves", len(board.MovesHistory())/2+1)
}

func LastMove(board chess.Board) Metric {
	movesCount := len(board.MovesHistory())
	if movesCount == 0 {
		return nil
	}

	return New("Last move", board.MovesHistory()[movesCount-1])
}

func Material(board chess.Board) Metric {
	callback := func(side chess.Side) int {
		var weight int
		for _, piece := range board.Squares().GetAllPieces(side) {
			weight += int(piece.Weight())
		}

		return weight
	}

	return New("Material value", []int{callback(chess.SideWhite), callback(chess.SideBlack)})
}

func MaterialDifference(board chess.Board) Metric {
	material := Material(board).Value().([]int)

	return New("Material advantage difference", material[0]-material[1])
}
