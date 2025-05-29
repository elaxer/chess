package resolver

import (
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
)

func UnresolveFrom(from, to position.Position, board chess.Board) (position.Position, error) {
	piece, err := board.Squares().GetByPosition(from)
	if err != nil {
		return from, err
	}

	hasSamePiece, hasSameFile, hasSameRank := false, false, false
	for _, samePiece := range board.Squares().GetPieces(piece.Notation(), piece.Side()) {
		samePiecePosition := board.Squares().GetByPiece(samePiece)
		if samePiecePosition == from || !board.LegalMoves(samePiece).ContainsOne(to) {
			continue
		}

		hasSamePiece = true
		hasSameFile = hasSameFile || samePiecePosition.File == from.File
		hasSameRank = hasSameRank || samePiecePosition.Rank == from.Rank
		if hasSameFile && hasSameRank {
			break
		}
	}

	unresolvedFrom := position.NewNull()
	if hasSameRank || (hasSamePiece && !hasSameFile) {
		unresolvedFrom.File = from.File
	}
	if hasSameFile {
		unresolvedFrom.Rank = from.Rank
	}

	return unresolvedFrom, nil
}
