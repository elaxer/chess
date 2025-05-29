package piece

import (
	"encoding/json"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
)

const (
	NotationQueen = "Q"
	WeightQueen   = 9
)

type Queen struct {
	*sliding
}

func NewQueen(side chess.Side) *Queen {
	return &Queen{&sliding{&abstract{side, false}}}
}

func (q *Queen) Side() chess.Side {
	return q.side
}

func (q *Queen) PseudoMoves(from position.Position, squares *chess.Squares) position.Set {
	moves := mapset.NewSetWithSize[position.Position](27)
	for _, direction := range allDirections {
		for move := range q.slide(from, direction, squares) {
			moves.Add(move)
		}
	}

	return moves
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
		"side":     q.abstract.side,
		"notation": q.Notation(),
	})
}
