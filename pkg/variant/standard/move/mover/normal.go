package mover

import (
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/variant/standard/move/move"
	"github.com/elaxer/chess/pkg/variant/standard/move/result"
)

type Normal struct {
}

func (m *Normal) Make(normal *move.Normal, board chess.Board) (*result.Normal, error) {
	if err := normal.Validate(); err != nil {
		return nil, err
	}

	pieceResult, err := movePiece(normal.Piece, normal.PieceNotation, board)
	if err != nil {
		return nil, err
	}

	pieceResult.Abstract = newAbstractResult(board)

	return &result.Normal{Piece: pieceResult, InputMove: *normal}, nil
}
