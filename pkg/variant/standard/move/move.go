package move

import (
	"errors"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/variant/standard/move/resolver"
)

func FromNotation(notation string, board chess.Board) (chess.Move, error) {
	if move, err := NewCastling(notation); err == nil {
		return move, nil
	}

	if move, err := NewPromotion(notation); err == nil {
		from, err := resolver.ResolveFrom(move.From, move.To, chess.NotationPawn, board)
		if err != nil {
			return nil, err
		}

		move.From = from

		return move, nil
	}

	if move, err := NewNormal(notation); err == nil {
		from, err := resolver.ResolveFrom(move.From, move.To, move.PieceNotation, board)
		if err != nil {
			return nil, err
		}

		move.From = from

		return move, nil
	}

	//todo
	return nil, errors.New("22")
}
