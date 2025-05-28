package metric

import "github.com/elaxer/chess/pkg/chess"

var AllFuncs = []MetricFunc{
	Turn,
	OppositeState,
	MovesCount,
	LastMove,
	AdvantageDifference,
}

func Turn(board chess.Board) Metric {
	return New("Turn", board.Turn())
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

func AdvantageDifference(board chess.Board) Metric {
	advantageFunc := func(side chess.Side) int {
		var advantage int
		for _, piece := range board.Squares().GetAllPieces(side) {
			advantage += int(piece.Weight())
		}

		return advantage
	}

	return New("Advantage", advantageFunc(chess.SideWhite)-advantageFunc(chess.SideBlack))
}
