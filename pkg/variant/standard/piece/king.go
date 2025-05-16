package piece

import (
	"encoding/json"
	"math"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
)

const (
	NotationKing = "K"
	WeightKing   = math.MaxUint8
)

type King struct {
	*basePiece
}

func NewKing(side chess.Side) *King {
	return &King{&basePiece{side, false}}
}

func (k *King) Moves(board chess.Board) position.Set {
	pos := board.Squares().GetByPiece(k)
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

	moves := mapset.NewSetWithSize[position.Position](8)
	for _, move := range positions {
		if err := move.Validate(); err != nil {
			continue
		}

		if piece, err := board.Squares().GetByPosition(move); err != nil && k.canMove(piece, k.side) {
			moves.Add(move)
		}
	}

	return k.legalMoves(board, k, moves)
}

func (k *King) Notation() string {
	return NotationKing
}

func (k *King) Weight() uint8 {
	return WeightKing
}

func (k *King) String() string {
	if k.side == chess.SideBlack {
		return "k"
	}

	return "K"
}

func (k *King) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"side":     k.side,
		"notation": k.Notation(),
	})
}
