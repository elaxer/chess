package piece

import (
	"encoding/json"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/set"
)

type Bishop struct {
	*sliding
}

func NewBishop(side chess.Side) *Bishop {
	return &Bishop{&sliding{&basePiece{side, false}}}
}

func (b *Bishop) Moves(board chess.Board) *position.Set {
	pos := board.Squares().GetByPiece(b).Position
	directions := [4]position.Position{
		position.New(1, 1),   // Diagonal up-right
		position.New(-1, -1), // Diagonal down-left
		position.New(1, -1),  // Diagonal down-right
		position.New(-1, 1),  // Diagonal up-left
	}

	moves := set.FromSlice(make([]position.Position, 0, 13))
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

func (b *Bishop) Notation() chess.PieceNotation {
	return chess.NotationBishop
}

func (b *Bishop) Weight() uint8 {
	return chess.WeightBishop
}

func (b *Bishop) String() string {
	return string(b.Notation())
}

func (b *Bishop) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"side":     b.side,
		"notation": b.Notation(),
	})
}
