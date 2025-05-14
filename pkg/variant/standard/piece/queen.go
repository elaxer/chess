package piece

import (
	"encoding/json"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
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
	square := board.Squares().GetByPiece(q)

	square.SetPiece(q.Rook)
	moves := q.Rook.Moves(board)

	square.SetPiece(q.Bishop)
	moves = moves.Union(q.Bishop.Moves(board))

	square.SetPiece(q)

	return q.legalMoves(board, q, moves)
}

func (q *Queen) Notation() chess.PieceNotation {
	return chess.NotationQueen
}

func (q *Queen) Weight() uint8 {
	return chess.WeightQueen
}

func (q *Queen) String() string {
	return string(q.Notation())
}

func (q *Queen) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"side":     q.basePiece.side,
		"notation": q.Notation(),
	})
}
