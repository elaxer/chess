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
func ResolveFrom(move move.Piece, pieceNotation string, board chess.Board, turn chess.Side) (position.Position, error) {
	if move.From.IsFull() {
		return move.From, nil
	}
	if err := move.Validate(); err != nil {
		return position.NewEmpty(), err
	}

	pieces := make([]chess.Piece, 0, 8)
	for _, piece := range board.Squares().GetPieces(pieceNotation, turn) {
		if board.LegalMoves(piece).ContainsOne(move.To) {
			pieces = append(pieces, piece)
		}
	}

	if len(pieces) == 0 {
		return position.NewEmpty(), fmt.Errorf("%w: no moves found", Err)
	}
	if len(pieces) == 1 {
		return board.Squares().GetByPiece(pieces[0]), nil
	}

	resolvedFrom := move.From
	for _, piece := range pieces {
		pos := board.Squares().GetByPiece(piece)
		if move.From.Rank.IsNull() && pos.File == move.From.File {
			resolvedFrom.Rank = pos.Rank
		}
		if move.From.File.IsNull() && pos.Rank == move.From.Rank {
			resolvedFrom.File = pos.File
		}
	}

	return resolvedFrom, nil
}
