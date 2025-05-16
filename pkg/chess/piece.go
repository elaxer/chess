package chess

import "github.com/elaxer/chess/pkg/chess/position"

// Piece это интерфейс шахматной фигуры.
type Piece interface {
	// Side возвращает сторону фигуры.
	Side() Side
	// IsMoved возвращает true, если фигура ходила.
	IsMoved() bool
	// MarkMoved помечает, что фигура сделала ход.
	MarkMoved()
	// Moves возвращает все возможные ходы фигуры на доске.
	// todo PseudoMoves(fromPosition)
	Moves(board Board) position.Set
	// Notation возвращает нотацию фигуры.
	Notation() string
	// Weight возвращает вес фигуры.
	Weight() uint8
}
