package staterule

import (
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
)

func Check(board chess.Board, side chess.Side) chess.State {
	_, kingPosition := board.Squares().GetPiece(piece.NotationKing, side)

	if board.Moves(!side).ContainsOne(kingPosition) {
		return chess.StateCheck
	}

	return chess.StateClear
}
