package mover

import (
	"github.com/elaxer/chess/pkg/chess"

	"github.com/elaxer/chess/pkg/variant/standard/move/move"
	"github.com/elaxer/chess/pkg/variant/standard/move/resolver"
	"github.com/elaxer/chess/pkg/variant/standard/move/validator"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
)

// Promotion это структура, реализующая интерфейс Mover для выполнения и проверки ходов превращения фигур.
// Она отвечает за логику, связанную с превращением пешки в другую фигуру на шахматной доске.
type Promotion struct {
}

func (m *Promotion) Make(move *move.Promotion, board chess.Board) (chess.Move, error) {
	var err error
	move.Normal, err = resolver.ResolveNormal(move.Normal, board, board.Turn())
	if err != nil {
		return nil, err
	}
	if err := validator.ValidatePromotion(move, board); err != nil {
		return nil, err
	}

	unresolvedFrom, err := resolver.UnresolveFrom(move.From, move.To, board)
	if err != nil {
		return nil, err
	}

	capturedPiece, err := board.Squares().MovePiece(move.From, move.To)
	if err != nil {
		return nil, err
	}

	piece := piece.New(move.NewPieceNotation, board.Turn())
	piece.MarkMoved()

	board.Squares().PlacePiece(piece, move.To)

	modifyNormal(move.Normal, capturedPiece, board)
	move.From = unresolvedFrom

	return move, nil
}
