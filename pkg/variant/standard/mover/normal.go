package mover

import (
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/variant/standard/move"
	resolver "github.com/elaxer/chess/pkg/variant/standard/moveresolver"
	validator "github.com/elaxer/chess/pkg/variant/standard/movevalidator"
	"github.com/elaxer/chess/pkg/variant/standard/state"
)

type Normal struct {
}

func (m *Normal) Make(move *move.Normal, board chess.Board) (chess.Move, error) {
	move, err := resolver.ResolveNormal(move, board, board.Turn())
	if err != nil {
		return nil, err
	}

	if err := validator.ValidateNormal(move, board); err != nil {
		return nil, err
	}

	capturedPiece, err := board.Squares().MovePiece(move.From, move.To)
	if err != nil {
		return nil, err
	}

	piece, _ := board.Squares().GetByPosition(move.To)
	piece.MarkMoved()

	modifyNormal(move, capturedPiece, board)

	return move, nil
}

func modifyNormal(move *move.Normal, capturedPiece chess.Piece, board chess.Board) {
	move.IsCapture = capturedPiece != nil
	move.CapturedPiece = capturedPiece
	move.CheckMate.IsCheck = board.State(!board.Turn()) == state.Check
	move.CheckMate.IsMate = board.State(!board.Turn()) == state.Mate
}
