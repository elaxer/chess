package formatter

import (
	"strconv"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/variant/standard/move"
)

func FormatNormalMove(move *move.Normal, board chess.Board) (chess.RawMove, error) {
	if err := move.Validate(); err != nil {
		return "", err
	}

	square := board.Squares().GetByPosition(move.From)
	piece := square.Piece

	hasSameFile, hasSameRank := false, false
	for _, samePiece := range board.Squares().GetPieces(piece.Notation(), piece.Side()) {
		samePiecePosition := board.Squares().GetByPiece(samePiece).Position
		if samePiecePosition == square.Position {
			continue
		}
		if samePiecePosition.File != square.Position.File && samePiecePosition.Rank != square.Position.Rank {
			continue
		}
		if !samePiece.Moves(board).Has(move.To) {
			continue
		}

		if samePiecePosition.File == square.Position.File {
			hasSameFile = true
		}
		if samePiecePosition.Rank == square.Position.Rank {
			hasSameRank = true
		}

		if hasSameFile && hasSameRank {
			break
		}
	}

	notation := string(move.PieceNotation)

	if hasSameRank {
		notation += square.Position.File.String()
	}
	if hasSameFile {
		notation += strconv.Itoa(int(square.Position.Rank))
	}

	if move.IsCapture {
		notation += "x"
	}
	notation += move.To.String()
	if move.IsCheck {
		notation += "+"
	}
	if move.IsMate {
		notation += "#"
	}

	return chess.RawMove(notation), nil
}
