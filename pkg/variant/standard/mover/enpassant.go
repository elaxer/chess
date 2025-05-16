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
	p, err := board.Squares().GetByPosition(epMove.From)
	if err != nil {
		return nil, err
	}

	pawn, ok := p.(*piece.Pawn)
	if !ok {
		return nil, err
	}

	if !pawn.Moves(board).ContainsOne(epMove.To) {
		return nil, err
	}

	_, err = board.Squares().GetByPosition(epMove.To)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
