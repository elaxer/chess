package staterule

import "github.com/elaxer/chess/pkg/chess"

func Check(board chess.Board, side chess.Side) chess.State {
	_, kingPosition := board.Squares().GetPiece(chess.NotationKing, side)

	if board.Moves(!side).ContainsOne(kingPosition) {
		return chess.StateCheck
	}

	return chess.StateClear
}
