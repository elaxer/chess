package resolver

import (
	"errors"
	"fmt"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/move"
)

var Err = errors.New("ошибка резолвинга")

// ResolveNormal определяет стартовую позицию фигуры, которая будет ходить.
// from - данные о стартовой позиции фигуры. Они могут быть заполнены не полностью.
func ResolveNormal(move *move.Normal, board chess.Board) (*move.Normal, error) {
	if move.From.Validate() == nil {
		return move, nil
	}
	if err := move.To.Validate(); err != nil {
		return nil, err
	}

	pieces := make([]chess.Piece, 0, 8)
	for _, piece := range board.Squares().GetPieces(move.PieceNotation, board.Turn()) {
		if piece.Moves(board).Has(move.To) {
			pieces = append(pieces, piece)
		}
	}

	if len(pieces) == 0 {
		return nil, fmt.Errorf("%w: не найдено подходящих фигур для хода", Err)
	}
	if len(pieces) == 1 {
		move.From = board.Squares().GetByPiece(pieces[0]).Position

		return move, nil
	}

	for _, piece := range pieces {
		pos := board.Squares().GetByPiece(piece).Position
		if move.From.Rank == 0 && pos.File == move.From.File {
			move.From.Rank = pos.Rank
		}
		if move.From.File == 0 && pos.Rank == move.From.Rank {
			move.From.File = pos.File
		}
	}

	return move, nil
}

func UnresolveFrom(from, to position.Position, board chess.Board) (position.Position, error) {
	if err := from.Validate(); err != nil {
		return from, err
	}

	square := board.Squares().GetByPosition(from)
	hasSameFile, hasSameRank := false, false

	for _, samePiece := range board.Squares().GetPieces(square.Piece.Notation(), square.Piece.Side()) {
		samePiecePosition := board.Squares().GetByPiece(samePiece).Position
		if samePiecePosition == square.Position {
			continue
		}
		if !samePiece.Moves(board).Has(to) {
			continue
		}

		hasSameFile = hasSameFile || samePiecePosition.File == square.Position.File
		hasSameRank = hasSameRank || samePiecePosition.Rank == square.Position.Rank
		if hasSameFile && hasSameRank {
			break
		}
	}

	unresolvedFrom := position.Position{}
	if hasSameRank {
		unresolvedFrom.File = square.Position.File
	}
	if hasSameFile {
		unresolvedFrom.Rank = square.Position.Rank
	}

	return unresolvedFrom, nil
}
