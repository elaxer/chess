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
	moveCopy := *move
	normalCopy := *move.Normal
	moveCopy.Normal = &normalCopy

	moveCopy.Normal.From, err = resolver.ResolveFrom(moveCopy.Normal, board, board.Turn())
	if err != nil {
		return nil, err
	}

	if err := validator.ValidatePromotion(&moveCopy, board); err != nil {
		return nil, err
	}

	unresolvedFrom, err := resolver.UnresolveFrom(moveCopy.Normal, board)
	if err != nil {
		return nil, err
	}

	capturedPiece, err := board.Squares().MovePiece(moveCopy.From, moveCopy.To)
	if err != nil {
		return nil, err
	}

	piece := piece.New(moveCopy.NewPieceNotation, board.Turn())
	piece.MarkMoved()

	board.Squares().PlacePiece(piece, moveCopy.To)

	modifyNormal(moveCopy.Normal, capturedPiece, board)
	moveCopy.From = unresolvedFrom

	return move, nil
}
