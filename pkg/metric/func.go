package metric

import "github.com/elaxer/chess/pkg/chess"

var AllFuncs = []MetricFunc{
	OppositeState,
	MovesCount,
	LastMove,
	Material,
	MaterialDifference,
}

func OppositeState(board chess.Board) Metric {
	return New("Opposite state", board.State(!board.Turn()))
}

func MovesCount(board chess.Board) Metric {
	return New("Moves count", len(board.MovesHistory()))
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
