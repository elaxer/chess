package rule

import (
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
	"github.com/elaxer/chess/pkg/variant/standard/state/state"
)

func Check(board chess.Board, side chess.Side) chess.State {
	_, kingPosition := board.Squares().FindPiece(piece.NotationKing, side)
	if board.Moves(!side).ContainsOne(kingPosition) {
		return state.Check
	}

	return nil
}
