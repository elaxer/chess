package piece

import (
	"encoding/json"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/set"
	"github.com/elaxer/chess/pkg/variant/standard/move"
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

func (p *Pawn) Moves(board chess.Board) *position.Set {
	moves := p.movesForward(board).Union(p.movesDiagonal(board)).Union(p.movesEnPassant(board))

	return p.legalMoves(board, p, moves)
}

func (p *Pawn) Notation() chess.PieceNotation {
	return chess.NotationPawn
}

func (p *Pawn) Weight() uint8 {
	return chess.WeightPawn
}

// movesForward возвращает возможные ходы вперед для пешки.
// Пешка может двигаться на одну или две клетки вперед, если она не была перемещена ранее.
func (p *Pawn) movesForward(board chess.Board) *position.Set {
	direction := PawnRankDirection(p.side)
	pos := board.Squares().GetByPiece(p).Position

	moves := set.FromSlice(make([]position.Position, 0, 2))
	positions := [2]position.Position{
		position.New(pos.File, pos.Rank+direction*1),
		position.New(pos.File, pos.Rank+direction*2),
	}
	for i, move := range positions {
		square := board.Squares().GetByPosition(move)
		if (square == nil || !square.IsEmpty()) || (i == 1 && p.isMoved) {
			break
		}

		moves.Add(move)
	}

	return moves
}

// movesDiagonal возвращает возможные диагональные ходы для пешки.
// Пешка может бить противника по диагонали на одну клетку вперед.
// Если на диагонали нет противника, то возвращается пустой массив.
func (p *Pawn) movesDiagonal(board chess.Board) *position.Set {
	direction := PawnRankDirection(p.side)
	pos := board.Squares().GetByPiece(p).Position

	moves := set.FromSlice(make([]position.Position, 0, 2))
	positions := [2]position.Position{
		position.New(pos.File+1, pos.Rank+direction),
		position.New(pos.File-1, pos.Rank+direction),
	}
	for _, move := range positions {
		square := board.Squares().GetByPosition(move)
		if square != nil && !square.IsEmpty() && square.Piece.Side() != p.side {
			moves.Add(move)
		}
	}

	return moves
}

// todo descr and tests
func (p *Pawn) movesEnPassant(board chess.Board) *position.Set {
	moves := set.FromSlice(make([]position.Position, 0, 1))
	if p.side != board.Turn() {
		return moves
	}

	pos := board.Squares().GetByPiece(p).Position
	movesHistory := board.MovesHistory()
	movesHistoryCount := len(movesHistory)

	if movesHistoryCount < 3 {
		return moves
	}

	lastMove, ok := movesHistory[movesHistoryCount-1].(*move.Normal)
	if !ok || lastMove.PieceNotation != chess.NotationPawn {
		return moves
	}
	if lastMove.To.Rank-lastMove.From.Rank != 2*PawnRankDirection(!board.Turn()) {
		return moves
	}
	if lastMove.To.File+1 != pos.File || lastMove.To.File-1 != pos.File {
		return moves
	}

	moves.Add(position.New(lastMove.To.File, lastMove.From.Rank+PawnRankDirection(board.Turn())))

	return moves
}

func (p *Pawn) String() string {
	return string(p.Notation())
}

func (p *Pawn) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"side":     p.side,
		"notation": p.Notation(),
	})
}
