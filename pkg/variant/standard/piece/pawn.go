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
	*basePiece
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
	return &Pawn{&basePiece{side, false}}
}

func (p *Pawn) Moves(board chess.Board) position.Set {
	moves := p.movesForward(board).Union(p.movesDiagonal(board))

	return p.legalMoves(board, p, moves)
}

func (p *Pawn) Notation() string {
	return NotationPawn
}

func (p *Pawn) Weight() uint8 {
	return WeightPawn
}

// movesForward возвращает возможные ходы вперед для пешки.
// Пешка может двигаться на одну или две клетки вперед, если она не была перемещена ранее.
func (p *Pawn) movesForward(board chess.Board) position.Set {
	direction := PawnRankDirection(p.side)
	pos := board.Squares().GetByPiece(p)

	moves := mapset.NewSetWithSize[position.Position](2)
	positions := [2]position.Position{
		position.New(pos.File, pos.Rank+direction*1),
		position.New(pos.File, pos.Rank+direction*2),
	}
	for i, move := range positions {
		piece, err := board.Squares().GetByPosition(move)
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
func (p *Pawn) movesDiagonal(board chess.Board) position.Set {
	direction := PawnRankDirection(p.side)
	pos := board.Squares().GetByPiece(p)

	moves := mapset.NewSetWithSize[position.Position](2)
	positions := [2]position.Position{
		position.New(pos.File+1, pos.Rank+direction),
		position.New(pos.File-1, pos.Rank+direction),
	}
	for _, move := range positions {
		piece, err := board.Squares().GetByPosition(move)
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
