package chess

import "github.com/elaxer/chess/pkg/chess/position"

// Square это структура, представляющая клетку на шахматной доске.
// Содержит информацию о фигуре, находящейся в клетке и его позиции на доске.
type Square struct {
	// Position это позиция клетки на доске.
	Position position.Position `json:"position"`
	// Piece это фигура, находящаяся в клетке.
	// Если фигура отсутствует, то Piece будет nil.
	Piece Piece `json:"piece,omitempty"`
}

// SetPiece помещает фигуру в клетку.
// Если клетка уже занята другой фигурой, то она будет заменена.
// Чтобы удалить фигуру из клетки, нужно передать nil в качестве аргумента.
// Если передать nil, то клетка останется пустой.
func (s *Square) SetPiece(p Piece) {
	s.Piece = p
}

// IsEmpty проверяет, пуста ли клетка.
// Если фигура отсутствует, то клетка считается пустой.
func (s *Square) IsEmpty() bool {
	return s.Piece == nil
}
