package piece

import (
	"encoding/json"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
)

const (
	NotationBishop = "B"
	WeightBishop   = 3
)

type Bishop struct {
	*sliding
}

func NewBishop(side chess.Side) *Bishop {
	return &Bishop{&sliding{&basePiece{side, false}}}
}

func (b *Bishop) Moves(board chess.Board) position.Set {
	pos := board.Squares().GetByPiece(b)
	directions := [4]position.Position{
		position.New(1, 1),   // Diagonal up-right
		position.New(-1, -1), // Diagonal down-left
		position.New(1, -1),  // Diagonal down-right
		position.New(-1, 1),  // Diagonal up-left
	}

	moves := mapset.NewSetWithSize[position.Position](13)
	for _, direction := range directions {
		for i, j := pos.File+direction.File, pos.Rank+direction.Rank; b.isInRange(i, j); i, j = i+direction.File, j+direction.Rank {
			move := position.New(i, j)
			canMove, canContinue := b.slide(move, board)
			if canMove {
				moves.Add(move)
			}
			if !canContinue {
				break
			}
		}
	}

	return b.legalMoves(board, b, moves)
}

func (b *Bishop) Notation() string {
	return NotationBishop
}

func (b *Bishop) Weight() uint8 {
	return WeightBishop
}

func (b *Bishop) String() string {
	if b.side == chess.SideBlack {
		return "b"
	}

	return "B"
}

func (b *Bishop) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"side":     b.side,
		"notation": b.Notation(),
	})
}
