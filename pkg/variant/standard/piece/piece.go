package piece

import "github.com/elaxer/chess/pkg/chess"

// New создает новую фигуру по ее нотации и стороне.
// Если нотация не поддерживается, возвращает nil.
// Например, для создания ферзя нужно передать QueenNotation и сторону игрока.
func New(notation string, side chess.Side) chess.Piece {
	switch notation {
	case NotationPawn:
		return NewPawn(side)
	case NotationKing:
		return NewKing(side)
	case NotationQueen:
		return NewQueen(side)
	case NotationBishop:
		return NewBishop(side)
	case NotationKnight:
		return NewKnight(side)
	case NotationRook:
		return NewRook(side)
	}

	return nil
}
