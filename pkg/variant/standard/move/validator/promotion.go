package validator

import (
	"fmt"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/variant/standard/move"
)

var ErrPromotion = fmt.Errorf("%w: ошибка валидации хода с превращением пешки", Err)

func ValidatePromotion(move *move.Promotion, board chess.Board) error {
	if err := move.Validate(); err != nil {
		return err
	}
	if err := ValidateNormal(move.Normal, board); err != nil {
		return err
	}

	fromSquare := board.Squares().GetByPosition(move.From)
	if fromSquare == nil {
		return fmt.Errorf("%w: %w", ErrPromotion, chess.ErrSquareNotFound)
	}
	if fromSquare.Piece.Notation() != chess.NotationPawn {
		return fmt.Errorf("%w: превращение возможно только из пешки", ErrPromotion)
	}

	return nil
}
