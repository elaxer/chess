package mover

import (
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/move"
	"github.com/elaxer/chess/pkg/variant/standard/move/validator"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
)

// Promotion это структура, реализующая интерфейс Mover для выполнения и проверки ходов превращения фигур.
// Она отвечает за логику, связанную с превращением пешки в другую фигуру на шахматной доске.
type Promotion struct {
}

func (m *Promotion) Make(move *move.Promotion, board chess.Board) (chess.Move, error) {
	if err := validator.ValidatePromotion(move, board); err != nil {
		return nil, err
	}

	from := position.New(move.To.File, move.To.Rank-piece.PawnRankDirection(board.Turn()))

	board.MovePiece(from, move.To)

	board.Squares().GetByPosition(move.To).SetPiece(piece.New(move.NewPiece, board.Turn()))

	modifyCheckMate(move.CheckMate, board)

	return move, nil
}

func (m *Promotion) Undo(move chess.Move, board chess.Board) error {
	if err := move.Validate(); err != nil {
		return err
	}

	return nil
}
