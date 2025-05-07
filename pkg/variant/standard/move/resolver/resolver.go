package resolver

import (
	"errors"
	"fmt"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
)

var Err = errors.New("ошибка резолвинга")

// ResolveFrom определяет стартовую позицию фигуры, которая будет ходить.
// from - данные о стартовой позиции фигуры. Они могут быть заполнены не полностью.
func ResolveFrom(from, to position.Position, pieceNotation chess.PieceNotation, board chess.Board) (position.Position, error) {
	if from.Validate() == nil {
		return from, nil
	}
	if err := to.Validate(); err != nil {
		return from, err
	}

	pieces := make([]chess.Piece, 0, 8)
	for _, piece := range board.Squares().GetPieces(pieceNotation, board.Turn()) {
		if piece.Moves(board).Has(to) {
			pieces = append(pieces, piece)
		}
	}

	if len(pieces) == 0 {
		return from, fmt.Errorf("%w: не найдено подходящих фигур для хода", Err)
	}
	if len(pieces) == 1 {
		return board.Squares().GetByPiece(pieces[0]).Position, nil
	}

	for _, piece := range pieces {
		pos := board.Squares().GetByPiece(piece).Position
		if from.Rank == 0 && pos.File == from.File {
			from.Rank = pos.Rank
		}
		if from.File == 0 && pos.Rank == from.Rank {
			from.File = pos.File
		}
	}

	return from, nil
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
