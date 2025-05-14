package mover

import (
	"github.com/elaxer/chess/pkg/chess"
	mv "github.com/elaxer/chess/pkg/variant/standard/move"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
)

type EnPassant struct {
}

// todo
func (m *EnPassant) MakeMove(move chess.Move, board chess.Board) (chess.Move, error) {
	epMove, err := mv.NewNormal(move.Notation())
	if err != nil {
		return nil, err
	}

	/* валидация todo return err */
	fromSquare := board.Squares().GetByPosition(epMove.From)
	if fromSquare == nil {
		return nil, chess.ErrSquareNotFound
	}

	pawn, ok := fromSquare.Piece.(*piece.Pawn)
	if !ok {
		return nil, err
	}

	if !pawn.Moves(board).ContainsOne(epMove.To) {
		return nil, err
	}

	toSquare := board.Squares().GetByPosition(epMove.To)
	if toSquare == nil {
		return nil, chess.ErrSquareNotFound
	}

	return nil, nil
}
