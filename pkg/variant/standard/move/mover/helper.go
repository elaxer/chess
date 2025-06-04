package mover

import (
	"github.com/elaxer/chess/pkg/chess"
	mv "github.com/elaxer/chess/pkg/variant/standard/move/move"
	"github.com/elaxer/chess/pkg/variant/standard/move/resolver"
	"github.com/elaxer/chess/pkg/variant/standard/move/result"
	"github.com/elaxer/chess/pkg/variant/standard/move/validator"
)

func movePiece(move mv.Piece, movingPieceNotation string, board chess.Board) (result.Piece, error) {
	fullFrom, err := resolver.ResolveFrom(move, movingPieceNotation, board, board.Turn())
	if err != nil {
		return result.Piece{}, err
	}
	pieceMove := mv.NewPiece(fullFrom, move.To)

	shortenedFrom, err := resolver.UnresolveFrom(pieceMove, board)
	if err != nil {
		return result.Piece{}, err
	}

	if err := validator.ValidatePieceMove(pieceMove, movingPieceNotation, board); err != nil {
		return result.Piece{}, err
	}

	capturedPiece, err := board.Squares().MovePiece(pieceMove.From, pieceMove.To)
	if err != nil {
		return result.Piece{}, err
	}

	piece, err := board.Squares().FindByPosition(pieceMove.To)
	if err != nil {
		return result.Piece{}, err
	}

	piece.MarkMoved()

	return result.Piece{
		FromFull:      fullFrom,
		FromShortened: shortenedFrom,
		CapturedPiece: capturedPiece,
	}, nil
}

func newAbstractResult(board chess.Board) result.Abstract {
	return result.Abstract{
		MoveSide: board.Turn(),
		NewState: board.State(!board.Turn()),
	}
}
