package chess

import (
	"errors"

	"github.com/elaxer/chess/pkg/chess/position"
)

var ErrSquareNotFound = errors.New("не удалось найти клетку")

// todo сделать двумерным
type Squares []*Square

// AddPiece добавляет фигуру на клетку по ее позиции.
func (s Squares) AddPiece(piece Piece, pos position.Position) {
	s.GetByPosition(pos).SetPiece(piece)
}

// GetByPiece возвращает клетку по фигуре.
// Если клетка не найдена, возвращает nil.
func (s Squares) GetByPiece(piece Piece) *Square {
	for _, square := range s {
		if square.Piece == piece {
			return square
		}
	}

	return nil
}

// GetByPosition возвращает клетку по ее позиции.
// Если клетка не найдена, возвращает nil.
// Позиция должна быть в пределах доски.
func (s Squares) GetByPosition(pos position.Position) *Square {
	for _, square := range s {
		if square.Position == pos {
			return square
		}
	}

	return nil
}

// GetAllPiecesCount возвращает количество фигур для данной стороны.
func (s Squares) GetAllPiecesCount(side Side) int {
	return len(s.GetAllPieces(side))
}

// GetAllPieces возвращает все фигуры для данной стороны.
func (s Squares) GetAllPieces(side Side) []Piece {
	pieces := make([]Piece, 0, 16)

	for _, square := range s {
		if !square.IsEmpty() && square.Piece.Side() == side {
			pieces = append(pieces, square.Piece)
		}
	}

	return pieces
}

// GetPieces возвращает все фигуры определенного типа для указанной стороны.
// Например, GetPieces(KnightNotation, SideWhite) вернет всех белых коней.
// Если фигуры не найдены, вернет пустой массив.
func (s Squares) GetPieces(notation PieceNotation, side Side) []Piece {
	pieces := make([]Piece, 0, 8)
	for _, square := range s {
		if !square.IsEmpty() && square.Piece.Side() == side && square.Piece.Notation() == notation {
			pieces = append(pieces, square.Piece)
		}
	}

	return pieces
}

// GetPiece возвращает одну фигуру определенного типа для указанной стороны и его позицию.
// Если фигура не найдена, вернет nil.
func (s Squares) GetPiece(notation PieceNotation, side Side) (Piece, position.Position) {
	pieces := s.GetPieces(notation, side)
	if len(pieces) == 0 {
		return nil, position.Position{}
	}

	return pieces[0], s.GetByPiece(pieces[0]).Position
}

// GetAdvantage возвращает материальное преимущество стороны.
func (s Squares) GetAdvantage(side Side) uint8 {
	var advantage uint8
	for _, piece := range s.GetAllPieces(side) {
		advantage += piece.Weight()
	}

	return advantage - WeightKing
}

// GetAdvantageDifference возвращает разницу в материальном преимуществе между сторонами.
// Положительное значение означает преимущество белых.
// Отрицательное значение означает преимущество черных.
// Нулевое значение означает равенство.
func (s Squares) GetAdvantageDifference() int8 {
	return int8(s.GetAdvantage(SideWhite)) - int8(s.GetAdvantage(SideBlack))
}
