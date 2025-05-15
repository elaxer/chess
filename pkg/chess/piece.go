package chess

import "github.com/elaxer/chess/pkg/chess/position"

// todo перенести
const (
	NotationPawn   PieceNotation = ""
	NotationKnight PieceNotation = "N"
	NotationBishop PieceNotation = "B"
	NotationRook   PieceNotation = "R"
	NotationQueen  PieceNotation = "Q"
	NotationKing   PieceNotation = "K"
)

// Вес шахматных фигур
const (
	WeightPawn   = 1
	WeightKnight = 3
	WeightBishop = 3
	WeightRook   = 5
	WeightQueen  = 9
	WeightKing   = 255
)

// PieceNotation это тип, представляющий нотацию шахматной фигуры.
// Он используется для обозначения различных фигур в шахматах, таких как пешка, король, ферзь и т.д.
// Нотация пешки не имеет символа, поэтому она представлена пустой строкой.
// Например, нотация ферзя - "Q", нотация короля - "K", и т.д.
type PieceNotation string

// Piece это интерфейс шахматной фигуры.
type Piece interface {
	// Side возвращает сторону фигуры.
	Side() Side
	// IsMoved возвращает true, если фигура ходила.
	IsMoved() bool
	// SetMoved устанавливает, что фигура сделала ход.
	// todo переименовать в MarkMoved
	SetMoved()
	// Moves возвращает все возможные ходы фигуры на доске.
	// todo PseudoMoves(fromPosition)
	Moves(board Board) position.Set
	// Notation возвращает нотацию фигуры.
	Notation() PieceNotation
	// Weight возвращает вес фигуры.
	Weight() uint8
}
