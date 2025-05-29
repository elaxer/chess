package validator

import (
	"fmt"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/variant/standard/move/move"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
)

var ErrPromotion = fmt.Errorf("%w: promotion validation error", Err)

func ValidatePromotion(move *move.Promotion, board chess.Board) error {
	if err := move.Validate(); err != nil {
		return err
	}
	if err := ValidateNormal(move.Normal, board); err != nil {
		return err
	}

	p, err := board.Squares().GetByPosition(move.From)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrPromotion, err)
	}
	if p.Notation() != piece.NotationPawn {
		return fmt.Errorf("%w: promotion is only possible from a pawn", ErrPromotion)
	}

	return nil
}
