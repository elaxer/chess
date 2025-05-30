package resolver

import (
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/move/move"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
)

func UnresolveFrom(move *move.Normal, board chess.Board) (position.Position, error) {
	p, err := board.Squares().GetByPosition(move.From)
	if err != nil {
		return move.From, err
	}

	hasSamePiece, hasSameFile, hasSameRank := false, false, false
	for _, samePiece := range board.Squares().GetPieces(p.Notation(), p.Side()) {
		samePiecePosition := board.Squares().GetByPiece(samePiece)
		if samePiecePosition == move.From || !board.LegalMoves(samePiece).ContainsOne(move.To) {
			continue
		}

		hasSamePiece = true
		hasSameFile = hasSameFile || samePiecePosition.File == move.From.File
		hasSameRank = hasSameRank || samePiecePosition.Rank == move.From.Rank
		if hasSameFile && hasSameRank {
			break
		}
	}

	unresolvedFrom := position.NewNull()
	if hasSameRank || (hasSamePiece && !hasSameFile) || (!hasSamePiece && p.Notation() == piece.NotationPawn && move.IsCapture) {
		unresolvedFrom.File = move.From.File
	}
	if hasSameFile {
		unresolvedFrom.Rank = move.From.Rank
	}

	return unresolvedFrom, nil
}
