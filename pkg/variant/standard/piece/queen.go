package piece

import (
	"encoding/json"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
)

const (
	NotationQueen = "Q"
	WeightQueen   = 9
)

type Queen struct {
	*basePiece
	*Rook
	*Bishop
}

func NewQueen(side chess.Side) *Queen {
	return &Queen{&basePiece{side, false}, NewRook(side), NewBishop(side)}
}

func (q *Queen) Side() chess.Side {
	return q.side
}

func (q *Queen) Moves(board chess.Board) position.Set {
	pos := board.Squares().GetByPiece(q)

	board.Squares().AddPiece(q.Rook, pos)
	moves := q.Rook.Moves(board)

	board.Squares().AddPiece(q.Bishop, pos)
	moves = moves.Union(q.Bishop.Moves(board))

	board.Squares().AddPiece(q, pos)

	return q.legalMoves(board, q, moves)
}

func (q *Queen) Notation() string {
	return NotationQueen
}

func (q *Queen) Weight() uint8 {
	return WeightQueen
}

func (q *Queen) String() string {
	if q.side == chess.SideBlack {
		return "q"
	}

	return "Q"
}

func (q *Queen) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"side":     q.basePiece.side,
		"notation": q.Notation(),
	})
}
