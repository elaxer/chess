package rule

import (
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/variant/standard/state/state"
)

func Checkmate(board chess.Board, side chess.Side) chess.State {
	if Check(board, side) != nil && board.Moves(side).Cardinality() == 0 {
		return state.Checkmate
	}

	return nil
}
