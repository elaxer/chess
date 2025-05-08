package mover

import (
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/variant/standard/move"
	"github.com/elaxer/chess/pkg/variant/standard/move/resolver"
	"github.com/elaxer/chess/pkg/variant/standard/move/validator"
)

type Normal struct {
}

func (m *Normal) Make(move *move.Normal, board chess.Board) (chess.Move, error) {
	move, err := resolver.ResolveNormal(move, board)
	if err != nil {
		return nil, err
	}

	if err := validator.ValidateNormal(move, board); err != nil {
		return nil, err
	}

	capturedPiece := board.MovePiece(move.From, move.To)

	move.CapturedPiece = capturedPiece
	move.IsCapture = capturedPiece != nil
	modifyCheckMate(move.CheckMate, board)

	return move, nil
}
