package chess

import (
	"errors"
	"fmt"
	"iter"

	"github.com/elaxer/chess/pkg/chess/position"
)

var ErrSquareOutOfRange = errors.New("square position is out of board bounds")

type Squares struct {
	squares      [][]Piece
	edgePosition position.Position
}

func NewSquares(edgePosition position.Position) Squares {
	squares := make([][]Piece, edgePosition.Rank)
	for i := range squares {
		squares[i] = make([]Piece, edgePosition.File)
	}

	return Squares{squares, edgePosition}
}

func (s Squares) Items() [][]Piece {
	return s.squares
}

func (s Squares) Iter() iter.Seq2[position.Position, Piece] {
	return func(yield func(pos position.Position, piece Piece) bool) {
		for rank, row := range s.squares {
			for file, piece := range row {
				if !yield(position.New(position.File(file+1), position.Rank(rank+1)), piece) {
					return
				}
			}
		}
	}
}

func (s Squares) EdgePosition() position.Position {
	return s.edgePosition
}

// GetByPosition возвращает клетку по ее позиции.
// Если клетка не найдена, возвращает nil.
// Позиция должна быть в пределах доски.
func (s Squares) GetByPosition(position position.Position) (Piece, error) {
	if !position.IsInRange(s.edgePosition) {
		return nil, fmt.Errorf("%w (%s)", ErrSquareOutOfRange, position)
	}

	return s.squares[position.Rank-1][position.File-1], nil
}

// GetByPiece возвращает клетку по фигуре.
// Если клетка не найдена, возвращает nil.
func (s Squares) GetByPiece(piece Piece) position.Position {
	for rank, row := range s.squares {
		for file := range row {
			if row[file] == piece {
				return position.New(position.File(file+1), position.Rank(rank+1))
			}
		}
	}

	return position.Position{}
}

// GetAllPiecesCount возвращает количество фигур для данной стороны.
func (s Squares) GetAllPiecesCount(side Side) int {
	return len(s.GetAllPieces(side))
}

// GetAllPieces возвращает все фигуры для данной стороны.
func (s Squares) GetAllPieces(side Side) []Piece {
	pieces := make([]Piece, 0, 16)
	for _, row := range s.squares {
		for _, piece := range row {
			if piece != nil && piece.Side() == side {
				pieces = append(pieces, piece)
			}
		}
	}

	return pieces
}

// GetPieces возвращает все фигуры определенного типа для указанной стороны.
// Например, GetPieces(KnightNotation, SideWhite) вернет всех белых коней.
// Если фигуры не найдены, вернет пустой массив.
func (s Squares) GetPieces(notation string, side Side) []Piece {
	pieces := make([]Piece, 0, 8)
	for _, row := range s.squares {
		for _, piece := range row {
			if piece != nil && piece.Side() == side && piece.Notation() == notation {
				pieces = append(pieces, piece)
			}
		}
	}

	return pieces
}

// GetPiece возвращает одну фигуру определенного типа для указанной стороны и его позицию.
// Если фигура не найдена, вернет nil.
func (s Squares) GetPiece(notation string, side Side) (Piece, position.Position) {
	pieces := s.GetPieces(notation, side)
	if len(pieces) == 0 {
		return nil, position.Position{}
	}

	return pieces[0], s.GetByPiece(pieces[0])
}

// AddPiece добавляет фигуру на клетку по ее позиции.
func (s Squares) AddPiece(piece Piece, position position.Position) error {
	if !position.IsInRange(s.edgePosition) {
		return ErrSquareOutOfRange
	}

	s.squares[position.Rank-1][position.File-1] = piece

	return nil
}
