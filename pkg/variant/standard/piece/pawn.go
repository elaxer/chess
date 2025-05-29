package piece

import (
	"encoding/json"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
)

const (
	NotationPawn = ""
	WeightPawn   = 1
)

type Pawn struct {
	*abstract
}

// PawnRankDirection возвращает направление движения пешки для указанной стороны.
// Для белых движение будет идти вверх (+1), для черных - вниз (-1).
func PawnRankDirection(side chess.Side) position.Rank {
	if side == chess.SideBlack {
		return -1
	}

	return 1
}

func NewPawn(side chess.Side) *Pawn {
	return &Pawn{&abstract{side, false}}
}

func (p *Pawn) PseudoMoves(from position.Position, squares *chess.Squares) position.Set {
	direction := PawnRankDirection(p.side)

	return p.movesForward(from, direction, squares).Union(p.movesDiagonal(from, direction, squares))
}

func (p *Pawn) Notation() string {
	return NotationPawn
}

func (p *Pawn) Weight() uint8 {
	return WeightPawn
}

// movesForward возвращает возможные ходы вперед для пешки.
// Пешка может двигаться на одну или две клетки вперед, если она не была перемещена ранее.
func (p *Pawn) movesForward(from position.Position, direction position.Rank, squares *chess.Squares) position.Set {
	moves := mapset.NewSetWithSize[position.Position](2)
	positions := [2]position.Position{
		position.New(from.File, from.Rank+direction*1),
		position.New(from.File, from.Rank+direction*2),
	}
	for i, move := range positions {
		piece, err := squares.GetByPosition(move)
		if (err != nil || piece != nil) || (i == 1 && p.isMoved) {
			break
		}

		moves.Add(move)
	}

	return moves
}

// movesDiagonal возвращает возможные диагональные ходы для пешки.
// Пешка может бить противника по диагонали на одну клетку вперед.
// Если на диагонали нет противника, то возвращается пустой массив.
func (p *Pawn) movesDiagonal(from position.Position, direction position.Rank, squares *chess.Squares) position.Set {
	moves := mapset.NewSetWithSize[position.Position](2)
	positions := [2]position.Position{
		position.New(from.File+1, from.Rank+direction),
		position.New(from.File-1, from.Rank+direction),
	}
	for _, move := range positions {
		piece, err := squares.GetByPosition(move)
		if err == nil && piece != nil && piece.Side() != p.side {
			moves.Add(move)
		}
	}

	return moves
}

func (p *Pawn) String() string {
	if p.side == chess.SideBlack {
		return "p"
	}

	return "P"
}

func (p *Pawn) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"side":     p.side,
		"notation": p.Notation(),
	})
}
