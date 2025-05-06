package piece

import (
	"encoding/json"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/set"
)

type King struct {
	*basePiece
}

func NewKing(side chess.Side) *King {
	return &King{&basePiece{side, false}}
}

func (k *King) Moves(board chess.Board) *position.Set {
	pos := board.Squares().GetByPiece(k).Position
	positions := [8]position.Position{
		position.New(pos.File, pos.Rank+1),
		position.New(pos.File, pos.Rank-1),
		position.New(pos.File+1, pos.Rank),
		position.New(pos.File-1, pos.Rank),
		position.New(pos.File+1, pos.Rank+1),
		position.New(pos.File-1, pos.Rank-1),
		position.New(pos.File+1, pos.Rank-1),
		position.New(pos.File-1, pos.Rank+1),
	}

	moves := set.FromSlice(make([]position.Position, 0, 8))
	for _, move := range positions {
		if err := move.Validate(); err != nil {
			continue
		}

		if square := board.Squares().GetByPosition(move); k.canMove(square, k.side) {
			moves.Add(move)
		}
	}

	return k.legalMoves(board, k, moves)
}

func (k *King) Notation() chess.PieceNotation {
	return chess.NotationKing
}

func (k *King) Weight() uint8 {
	return chess.WeightKing
}

func (k *King) String() string {
	return string(k.Notation())
}

func (k *King) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"side":     k.side,
		"notation": k.Notation(),
	})
}
