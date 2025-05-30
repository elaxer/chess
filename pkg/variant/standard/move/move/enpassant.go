package move

import (
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
)

type EnPassant struct {
	*Normal
}

func NewEnPassant(from, to position.Position, capturedPiece chess.Piece, checkMate CheckMate) *EnPassant {
	return &EnPassant{
		Normal: &Normal{
			PieceNotation: piece.NotationPawn,
			From:          from,
			To:            to,
			IsCapture:     true,
			CapturedPiece: capturedPiece,
			CheckMate:     checkMate,
		},
	}
}
