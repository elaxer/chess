package rule

import (
	"slices"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/variant/standard/move/move"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
	"github.com/elaxer/chess/pkg/variant/standard/state/state"
)

func Stalemate(board chess.Board, side chess.Side) chess.State {
	if Check(board, side) == nil && board.Moves(side).Cardinality() == 0 {
		return state.Stalemate
	}

	return nil
}

func FiftyMove(board chess.Board, side chess.Side) chess.State {
	moves := slices.Clone(board.MovesHistory())
	slices.Reverse(moves)

	count := 0
	for _, m := range moves {
		normalMove, ok := m.(*move.Normal)
		if !ok || normalMove.PieceNotation == piece.NotationPawn || normalMove.IsCapture {
			count = 0

			continue
		}

		count++
	}

	if count/2+1 >= 50 {
		return state.DrawFiftyMoves
	}

	return nil
}
