package chess

import "github.com/elaxer/chess/pkg/chess/position"

// Board интерфейс описывает основные методы для работы с шахматной доской.
type Board interface {
	// Square возвращает клетки на доске.
	// Содержит все клетки на доске и предоставляет методы для работы с ними.
	Squares() Squares
	// Turn возвращает текущую сторону, чей ход.
	// Например, если текущая сторона - белые, то функция вернет SideWhite.
	Turn() Side
	// MovesHistory возвращает список сыгранных ходов.
	MovesHistory() []Move
	// Moves возвращает все доступные ходы для указанной стороны.
	Moves(side Side) *position.Set
	// State возвращает текущее состояние доски.
	// Например, если на доске шах, то функция вернет StateCheck.
	State() State

	BoardMover
}

// BoardMover это интерфейс дающий методы управлением доски
type BoardMover interface {
	// MakeMove выполняет ход на доске.
	// Возвращает ошибку в случае невозможности хода.
	// Меняет очередь хода.
	MakeMove(move Move) error
	// NextTurn переключает ход на следующую сторону.
	// Например, если текущая сторона - белые, то после вызова функции
	// текущей стороной будут черные.
	NextTurn()
	// MovePiece перемещает фигуру с одной клетки на другую.
	// Возвращает съеденную фигуру. Если фигуры не было съедено - возвращает nil.
	MovePiece(from, to position.Position) (capturedPiece Piece)
}
