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

func (p *basePiece) MarkMoved() {
	p.isMoved = true
}

// legalMoves фильтрует возможные ходы фигуры, исключая те, которые ставят короля под шах.
// Если фигура не принадлежит текущей стороне хода, то фильтрации происходить не будет.
func (p *basePiece) legalMoves(board chess.Board, piece chess.Piece, moves position.Set) position.Set {
	if piece.Side() != board.Turn() {
		return moves
	}

	fromPosition := board.Squares().GetByPiece(piece)

	legalMoves := mapset.NewSetWithSize[position.Position](moves.Cardinality())
	for move := range moves.Iter() {
		p.temporaryMoving(fromPosition, move, board.Squares(), func() {
			_, kingPosition := board.Squares().GetPiece(NotationKing, board.Turn())
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
func (b *basePiece) temporaryMoving(fromPosition, toPosition position.Position, squares chess.Squares, callback func()) {
	movingPiece, _ := squares.GetByPosition(fromPosition)
	toSquare, _ := squares.GetByPosition(toPosition)

	toSquare, movingPiece = movingPiece, nil
	defer func() {
		toSquare, movingPiece = movingPiece, toSquare
	}()

	callback()
}

// canMove проверяет, может ли фигура переместиться на указанную клетку.
// Если клетка существует и пуста или занята фигурой противника, то перемещение возможно.
// Метод не должен использоваться для проверки ходов пешек.
func (p *basePiece) canMove(piece chess.Piece, pieceSide chess.Side) bool {
	return piece == nil || pieceSide != piece.Side()
}
