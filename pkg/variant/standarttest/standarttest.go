package standarttest

import (
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard"
)

type Placement struct {
	Piece    chess.Piece
	Position position.Position
}

func NewEmptyBoard(turn chess.Side, placements []Placement) chess.Board {
	board := standard.NewBoardFactory().CreateEmpty()
	if turn.IsBlack() {
		board.NextTurn()
	}
	for _, placement := range placements {
		board.Squares().AddPiece(placement.Piece, placement.Position)
	}

	return board
}
