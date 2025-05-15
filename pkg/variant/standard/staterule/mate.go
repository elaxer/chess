package staterule

import "github.com/elaxer/chess/pkg/chess"

func Mate(board chess.Board, side chess.Side) chess.State {
	if Check(board, side) != chess.StateClear && board.Moves(side).Cardinality() == 0 {
		return chess.StateMate
	}

	return chess.StateClear
}
