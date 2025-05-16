package piece

import (
	"encoding/json"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
)

const (
	NotationKnight = "N"
	WeightKnight   = 3
)

type Knight struct {
	*basePiece
}

func NewKnight(side chess.Side) *Knight {
	return &Knight{&basePiece{side, false}}
}

func (k *Knight) Moves(board chess.Board) position.Set {
	pos := board.Squares().GetByPiece(k)
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

	moves := mapset.NewSetWithSize[position.Position](8)
	for _, move := range positions {
		if piece, err := board.Squares().GetByPosition(move); err != nil && k.canMove(piece, k.side) {
			moves.Add(move)
		}
	}

	return k.legalMoves(board, k, moves)
}
func (k *Knight) Notation() string {
	return NotationKnight
}

func (k *Knight) Weight() uint8 {
	return WeightKnight
}

func (k *Knight) String() string {
	if k.side == chess.SideBlack {
		return "n"
	}

	return "N"
}

func (k *Knight) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"side":     k.side,
		"notation": k.Notation(),
	})
}
