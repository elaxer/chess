package staterule

import (
	"slices"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/variant/standard/move"
)

func Stalemate(board chess.Board, side chess.Side) chess.State {
	if Check(board, side) == chess.StateClear && board.Moves(side).Cardinality() == 0 {
		return chess.StateStalemate
	}

	return chess.StateClear
}

// todo
func ThreefoldRepetition(board chess.Board, side chess.Side) chess.State {
	return chess.StateClear
}

func FiftyMove(board chess.Board, side chess.Side) chess.State {
	moves := slices.Clone(board.MovesHistory())
	slices.Reverse(moves)

	count := 0
	for _, m := range moves {
		normalMove, ok := m.(*move.Normal)
		if !ok || normalMove.PieceNotation == chess.NotationPawn || normalMove.IsCapture {
			count = 0

			continue
		}

		count++
	}

	if count/2+1 >= 50 {
		return chess.StateDraw
	}

	return chess.StateClear
}
