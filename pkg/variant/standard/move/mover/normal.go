package mover

import (
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/variant/standard/move/move"
	"github.com/elaxer/chess/pkg/variant/standard/move/resolver"
	"github.com/elaxer/chess/pkg/variant/standard/move/validator"
)

type Normal struct {
}

func (m *Normal) Make(move *move.Normal, board chess.Board) (chess.Move, error) {
	var err error
	moveCopy := *move

	moveCopy.From, err = resolver.ResolveFrom(move, board, board.Turn())
	if err != nil {
		return nil, err
	}

	if err := validator.ValidateNormal(&moveCopy, board); err != nil {
		return nil, err
	}

	unresolvedFrom, err := resolver.UnresolveFrom(&moveCopy, board)
	if err != nil {
		return nil, err
	}

	capturedPiece, err := board.Squares().MovePiece(moveCopy.From, moveCopy.To)
	if err != nil {
		return nil, err
	}

	piece, _ := board.Squares().GetByPosition(moveCopy.To)
	piece.MarkMoved()

	modifyNormal(&moveCopy, capturedPiece, board)

	moveCopy.From = unresolvedFrom

	return &moveCopy, nil
}

func modifyNormal(move *move.Normal, capturedPiece chess.Piece, board chess.Board) {
	move.IsCapture = capturedPiece != nil
	move.CapturedPiece = capturedPiece
	move.NewBoardState = board.State(!board.Turn())
}
