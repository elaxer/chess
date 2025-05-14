package piece

import (
	"encoding/json"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
)

type Rook struct {
	*sliding
}

func NewRook(side chess.Side) *Rook {
	return &Rook{&sliding{&basePiece{side, false}}}
}

func (r *Rook) Moves(board chess.Board) position.Set {
	pos := board.Squares().GetByPiece(r).Position
	directions := [4]position.Position{
		position.New(1, 0),  // Right
		position.New(-1, 0), // Left
		position.New(0, 1),  // Up
		position.New(0, -1), // Down
	}

	moves := mapset.NewSetWithSize[position.Position](14)
	for _, direction := range directions {
		for i, j := pos.File+direction.File, pos.Rank+direction.Rank; r.isInRange(i, j); i, j = i+direction.File, j+direction.Rank {
			move := position.New(i, j)
			canMove, canContinue := r.slide(move, board)
			if canMove {
				moves.Add(move)
			}
			if !canContinue {
				break
			}
		}
	}

	return r.legalMoves(board, r, moves)
}

func (r *Rook) Notation() chess.PieceNotation {
	return chess.NotationRook
}

func (r *Rook) Weight() uint8 {
	return chess.WeightRook
}

func (r *Rook) String() string {
	return string(r.Notation())
}

func (r *Rook) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"side":     r.Side(),
		"notation": r.Notation(),
	})
}
