package staterule

import (
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
	"github.com/elaxer/chess/pkg/variant/standard/state"
)

func Check(board chess.Board, side chess.Side) chess.State {
	_, kingPosition := board.Squares().GetPiece(piece.NotationKing, side)

	if board.Moves(!side).ContainsOne(kingPosition) {
		return state.Check
	}

	return nil
}

func Mate(board chess.Board, side chess.Side) chess.State {
	if Check(board, side) != nil && board.Moves(side).Cardinality() == 0 {
		return state.Mate
	}

	return nil
}
