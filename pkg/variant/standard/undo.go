package standard

import (
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
)

type Undo struct {
	CapturedPiece     chess.Piece
	FromPieceWasMoved bool
	From, To          position.Position
}
