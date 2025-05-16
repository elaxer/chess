package piece

import (
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
)

// sliding - структура для фигур, которые могут двигаться по диагонали или вертикали/горизонтали
// (слон, ферзь, ладья).
// Она содержит базовую структуру фигуры и методы для проверки возможности движения.
type sliding struct {
	*basePiece
}

// slide - метод, который проверяет возможность движения фигуры по диагонали или вертикали/горизонтали.
// canMove определяет, может ли фигура переместиться на указанную позицию,
// canContinue определяет, может ли фигура продолжать движение в том же направлении.
func (s *sliding) slide(position position.Position, board chess.Board) (canMove bool, canContinue bool) {
	piece, err := board.Squares().GetByPosition(position)

	return s.canMove(piece, s.side), err == nil && piece == nil
}

// todo
func (s *sliding) isInRange(file position.File, rank position.Rank) bool {
	return position.New(file, rank).Validate() == nil
}
