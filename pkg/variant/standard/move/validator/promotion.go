package validator

import (
	"fmt"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/move"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
)

var ErrPromotion = fmt.Errorf("%w: ошибка валидации хода с превращением пешки", Err)

func ValidatePromotion(move *move.Promotion, board chess.Board) error {
	if err := move.Validate(); err != nil {
		return err
	}

	pos := position.New(move.To.File, move.To.Rank-piece.PawnRankDirection(board.Turn()))

	fromSquare := board.Squares().GetByPosition(pos)
	if fromSquare.IsEmpty() {
		return ErrEmptySquare
	}
	if fromSquare.Piece.Notation() != chess.NotationPawn {
		return fmt.Errorf("%w: превращение возможно только из пешки", ErrPromotion)
	}

	toSquare := board.Squares().GetByPosition(move.To)
	if !toSquare.IsEmpty() {
		return fmt.Errorf("%w: невозможно превратить фигуру на занятой клетке", ErrPromotion)
	}

	return nil
}
