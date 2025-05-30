package resolver

import (
	"errors"
	"fmt"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/move/move"
)

var Err = errors.New("resolving error")

// ResolveFrom определяет стартовую позицию фигуры, которая будет ходить.
// from - данные о стартовой позиции фигуры. Они могут быть заполнены не полностью.
func ResolveFrom(move *move.Normal, board chess.Board, turn chess.Side) (position.Position, error) {
	if move.From.Validate() == nil {
		return move.From, nil
	}
	if err := move.To.Validate(); err != nil {
		return position.NewNull(), err
	}

	pieces := make([]chess.Piece, 0, 8)
	for _, piece := range board.Squares().GetPieces(move.PieceNotation, turn) {
		if board.LegalMoves(piece).ContainsOne(move.To) {
			pieces = append(pieces, piece)
		}
	}

	if len(pieces) == 0 {
		return position.NewNull(), fmt.Errorf("%w: no moves found", Err)
	}
	if len(pieces) == 1 {
		return board.Squares().GetByPiece(pieces[0]), nil
	}

	resolvedFrom := move.From
	for _, piece := range pieces {
		pos := board.Squares().GetByPiece(piece)
		if move.From.Rank == 0 && pos.File == move.From.File {
			resolvedFrom.Rank = pos.Rank
		}
		if move.From.File == 0 && pos.Rank == move.From.Rank {
			resolvedFrom.File = pos.File
		}
	}

	return resolvedFrom, nil
}
