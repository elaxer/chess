package rule

import (
	"slices"

	"github.com/elaxer/chess/pkg/chess"
	mv "github.com/elaxer/chess/pkg/variant/standard/move/move"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
	"github.com/elaxer/chess/pkg/variant/standard/state/state"
)

func Stalemate(board chess.Board, side chess.Side) chess.State {
	if Check(board, side) == nil && board.Moves(side).Cardinality() == 0 {
		return state.Stalemate
	}

	return nil
}

func FiftyMoves(board chess.Board, side chess.Side) chess.State {
	moves := slices.Clone(board.MovesHistory())
	slices.Reverse(moves)

	count := 0
	for _, move := range moves {
		normalMove, ok := move.(*mv.Normal)
		if !ok || normalMove.PieceNotation == piece.NotationPawn || normalMove.IsCapture {
			count = 0
		} else {
			count++
		}
	}

	if count >= 50 {
		return state.DrawFiftyMoves
	}

	return nil
}
