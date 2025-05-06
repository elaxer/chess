package piece

import (
	"encoding/json"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/set"
)

type Knight struct {
	*basePiece
}

func NewKnight(side chess.Side) *Knight {
	return &Knight{&basePiece{side, false}}
}

func (k *Knight) Moves(board chess.Board) *position.Set {
	pos := board.Squares().GetByPiece(k).Position
	positions := [8]position.Position{
		position.New(pos.File+1, pos.Rank+2),
		position.New(pos.File-1, pos.Rank+2),
		position.New(pos.File+2, pos.Rank+1),
		position.New(pos.File-2, pos.Rank+1),
		position.New(pos.File-1, pos.Rank-2),
		position.New(pos.File-2, pos.Rank-1),
		position.New(pos.File+2, pos.Rank-1),
		position.New(pos.File+1, pos.Rank-2),
	}

	moves := set.FromSlice(make([]position.Position, 0, 8))
	for _, move := range positions {
		if square := board.Squares().GetByPosition(move); k.canMove(square, k.side) {
			moves.Add(move)
		}
	}

	return k.legalMoves(board, k, moves)
}
func (k *Knight) Notation() chess.PieceNotation {
	return chess.NotationKnight
}

func (k *Knight) Weight() uint8 {
	return chess.WeightKnight
}

func (k *Knight) String() string {
	return string(k.Notation())
}

func (k *Knight) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"side":     k.side,
		"notation": k.Notation(),
	})
}
