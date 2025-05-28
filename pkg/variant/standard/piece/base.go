package piece

import (
	"github.com/elaxer/chess/pkg/chess"
)

// base это базовая структура для шахматной фигуры.
// Она содержит базовые поля и вспомогательные методы для работы с фигурами.
type base struct {
	side    chess.Side
	isMoved bool
}

func (p *base) Side() chess.Side {
	return p.side
}

func (p *base) IsMoved() bool {
	return p.isMoved
}

func (p *base) MarkMoved() {
	p.isMoved = true
}

// canMove проверяет, может ли фигура переместиться на указанную клетку.
// Если клетка существует и пуста или занята фигурой противника, то перемещение возможно.
// Метод не должен использоваться для проверки ходов пешек.
func (p *base) canMove(squarePiece chess.Piece, pieceSide chess.Side) bool {
	return squarePiece == nil || pieceSide != squarePiece.Side()
}
