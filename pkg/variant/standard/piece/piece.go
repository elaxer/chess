package piece

import "github.com/elaxer/chess/pkg/chess"

// New создает новую фигуру по ее нотации и стороне.
// Если нотация не поддерживается, возвращает nil.
// Например, для создания ферзя нужно передать QueenNotation и сторону игрока.
func New(notation chess.PieceNotation, side chess.Side) chess.Piece {
	switch notation {
	case chess.NotationPawn:
		return NewPawn(side)
	case chess.NotationKing:
		return NewKing(side)
	case chess.NotationQueen:
		return NewQueen(side)
	case chess.NotationBishop:
		return NewBishop(side)
	case chess.NotationKnight:
		return NewKnight(side)
	case chess.NotationRook:
		return NewRook(side)
	}

	return nil
}
