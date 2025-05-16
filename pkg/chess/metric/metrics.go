package metric

import "github.com/elaxer/chess/pkg/chess"

var AllFuncs = []MetricFunc{
	Turn,
	State,
	OppositeState,
	MovesCount,
	LastMove,
	AdvantageDifference,
}

var stateStrings = map[chess.State]string{
	chess.StateClear:     "clear",
	chess.StateCheck:     "check",
	chess.StateDraw:      "draw",
	chess.StateMate:      "mate",
	chess.StateStalemate: "stalemate",
}

func Turn(board chess.Board) Metric {
	return New("Turn", board.Turn())
}

func State(board chess.Board) Metric {
	return New("State", stateStrings[board.State(board.Turn())])
}

func OppositeState(board chess.Board) Metric {
	return New("Opposite state", stateStrings[board.State(!board.Turn())])
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
