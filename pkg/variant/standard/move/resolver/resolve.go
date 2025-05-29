package resolver

import (
	"errors"
	"fmt"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/variant/standard/move/move"
)

var Err = errors.New("resolving error")

// ResolveNormal определяет стартовую позицию фигуры, которая будет ходить.
// from - данные о стартовой позиции фигуры. Они могут быть заполнены не полностью.
func ResolveNormal(move *move.Normal, board chess.Board, turn chess.Side) (*move.Normal, error) {
	if move.From.Validate() == nil {
		return move, nil
	}
	if err := move.To.Validate(); err != nil {
		return nil, err
	}

	pieces := make([]chess.Piece, 0, 8)
	for _, piece := range board.Squares().GetPieces(move.PieceNotation, turn) {
		if board.LegalMoves(piece).ContainsOne(move.To) {
			pieces = append(pieces, piece)
		}
	}

	if len(pieces) == 0 {
		return nil, fmt.Errorf("%w: no moves found", Err)
	}
	if len(pieces) == 1 {
		move.From = board.Squares().GetByPiece(pieces[0])

		return move, nil
	}

	for _, piece := range pieces {
		pos := board.Squares().GetByPiece(piece)
		if move.From.Rank == 0 && pos.File == move.From.File {
			move.From.Rank = pos.Rank
		}
		if move.From.File == 0 && pos.Rank == move.From.Rank {
			move.From.File = pos.File
		}
	}

	return move, nil
}
