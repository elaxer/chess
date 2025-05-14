package piece

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
)

// basePiece это базовая структура для шахматной фигуры.
// Она содержит базовые поля и вспомогательные методы для работы с фигурами.
type basePiece struct {
	side    chess.Side
	isMoved bool
}

func (p *basePiece) Side() chess.Side {
	return p.side
}

func (p *basePiece) IsMoved() bool {
	return p.isMoved
}

func (p *basePiece) SetMoved() {
	p.isMoved = true
}

// legalMoves фильтрует возможные ходы фигуры, исключая те, которые ставят короля под шах.
// Если фигура не принадлежит текущей стороне хода, то фильтрации происходить не будет.
func (p *basePiece) legalMoves(board chess.Board, piece chess.Piece, moves position.Set) position.Set {
	if piece.Side() != board.Turn() {
		return moves
	}

	fromSquare := board.Squares().GetByPiece(piece)

	legalMoves := mapset.NewSetWithSize[position.Position](moves.Cardinality())
	for move := range moves.Iter() {
		p.temporaryMoving(fromSquare, board.Squares().GetByPosition(move), func() {
			_, kingPosition := board.Squares().GetPiece(chess.NotationKing, board.Turn())
			if !board.Moves(!board.Turn()).ContainsOne(kingPosition) {
				legalMoves.Add(move)
			}
		})
	}

	return legalMoves
}

// temporaryMoving временно перемещает фигуру с одной клетки на другую и в контексте этого вызывает функцию.
// После выполнения функции, фигура возвращается на исходную позицию.
// Это используется для проверки возможности движения фигуры без изменения состояния доски.
func (b *basePiece) temporaryMoving(fromSquare, toSquare *chess.Square, callback func()) {
	fromSquarePiece := fromSquare.Piece
	toSquarePiece := toSquare.Piece

	fromSquare.SetPiece(nil)
	toSquare.SetPiece(fromSquarePiece)
	defer toSquare.SetPiece(toSquarePiece)
	defer fromSquare.SetPiece(fromSquarePiece)

	callback()
}

// canMove проверяет, может ли фигура переместиться на указанную клетку.
// Если клетка существует и пуста или занята фигурой противника, то перемещение возможно.
// Метод не должен использоваться для проверки ходов пешек.
func (p *basePiece) canMove(square *chess.Square, pieceSide chess.Side) bool {
	return square != nil && (square.IsEmpty() || pieceSide != square.Piece.Side())
}
