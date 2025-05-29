package validator

import (
	"fmt"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/variant/standard/move/move"
)

var ErrNormal = fmt.Errorf("%w: normal move validation error", Err)

func ValidateNormal(move *move.Normal, board chess.Board) error {
	if err := move.Validate(); err != nil {
		return err
	}

	piece, err := board.Squares().GetByPosition(move.From)
	if err != nil || piece == nil {
		return ErrEmptySquare
	}

	if piece.Side() != board.Turn() {
		return fmt.Errorf("%w: wrong side", ErrNormal)
	} else if !board.LegalMoves(piece).ContainsOne(move.To) {
		return fmt.Errorf("%w: piece doesn't have such move", ErrNormal)
	}

	return nil
}
