package validator

import (
	"fmt"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/variant/standard/move"
)

var ErrNormal = fmt.Errorf("%w: ошибка валидации обычного хода", Err)

func ValidateNormal(move *move.Normal, board chess.Board) error {
	if err := move.Validate(); err != nil {
		return err
	}

	fromSquare := board.Squares().GetByPosition(move.From)
	if fromSquare == nil || fromSquare.IsEmpty() {
		return ErrEmptySquare
	}

	if piece := fromSquare.Piece; piece.Side() != board.Turn() {
		return fmt.Errorf("%w: неверная сторона хода", ErrNormal)
	} else if !piece.Moves(board).ContainsOne(move.To) {
		return fmt.Errorf("%w: фигура не имеет такого хода", ErrNormal)
	}

	return nil
}
