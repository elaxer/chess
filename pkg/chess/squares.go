package chess

import (
	"errors"
	"fmt"
	"iter"
	"slices"

	"github.com/elaxer/chess/pkg/chess/position"
)

var MaxEdgePosition = position.New(position.FileMax, position.RankMax)

var ErrSquareOutOfRange = errors.New("square position is out of range")

type Squares struct {
	rows         [][]Piece
	edgePosition position.Position
}

func NewSquares(edgePosition position.Position) *Squares {
	if !edgePosition.IsBoundaries(MaxEdgePosition) {
		panic("cannot create squares with size greater than max size")
	}

	squares := make([][]Piece, edgePosition.Rank)
	for i := range squares {
		squares[i] = make([]Piece, edgePosition.File)
	}

	return &Squares{squares, edgePosition}
}

func (s *Squares) EdgePosition() position.Position {
	return s.edgePosition
}

func (s *Squares) Iter() iter.Seq2[position.Position, Piece] {
	return func(yield func(position position.Position, piece Piece) bool) {
		for rank, row := range s.rows {
			for file, piece := range row {
				if !yield(position.New(position.File(file+1), position.Rank(rank+1)), piece) {
					return
				}
			}
		}
	}
}

func (s *Squares) IterByRows(backward bool) iter.Seq2[position.Rank, iter.Seq2[position.File, Piece]] {
	rows := slices.Clone(s.rows)
	if backward {
		slices.Reverse(rows)
	}

	return func(yield func(position.Rank, iter.Seq2[position.File, Piece]) bool) {
		for rank, row := range rows {
			isContinue := yield(position.Rank(rank+1), func(yield func(position.File, Piece) bool) {
				for file, piece := range row {
					if !yield(position.File(file+1), piece) {
						return
					}
				}
			})

			if !isContinue {
				return
			}
		}
	}
}

func (s *Squares) IterByDirection(from, direction position.Position) iter.Seq2[position.Position, Piece] {
	return func(yield func(position position.Position, piece Piece) bool) {
		position := position.New(from.File+direction.File, from.Rank+direction.Rank)
		for position.IsBoundaries(s.edgePosition) {
			if !yield(position, s.rows[position.Rank-1][position.File-1]) {
				return
			}

			position.File += direction.File
			position.Rank += direction.Rank
		}
	}
}

// FindByPosition возвращает клетку по ее позиции.
// Если клетка не найдена, возвращает nil.
// Позиция должна быть в пределах доски.
func (s *Squares) FindByPosition(position position.Position) (Piece, error) {
	if !position.IsBoundaries(s.edgePosition) {
		return nil, fmt.Errorf("%w (%s)", ErrSquareOutOfRange, position)
	}

	return s.rows[position.Rank-1][position.File-1], nil
}

// GetByPiece возвращает клетку по фигуре.
// Если клетка не найдена, возвращает nil.
func (s *Squares) GetByPiece(piece Piece) position.Position {
	for rank, row := range s.rows {
		for file := range row {
			if row[file] == piece {
				return position.New(position.File(file+1), position.Rank(rank+1))
			}
		}
	}

	return position.NewEmpty()
}

// GetAllPieces возвращает все фигуры для данной стороны.
func (s *Squares) GetAllPieces(side Side) []Piece {
	pieces := make([]Piece, 0, 16)
	for _, row := range s.rows {
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
func (s *Squares) GetPieces(notation string, side Side) []Piece {
	pieces := make([]Piece, 0, 8)
	for _, row := range s.rows {
		for _, piece := range row {
			if piece != nil && piece.Side() == side && piece.Notation() == notation {
				pieces = append(pieces, piece)
			}
		}
	}

	return pieces
}

// FindPiece возвращает одну фигуру определенного типа для указанной стороны и его позицию.
// Если фигура не найдена, вернет nil.
func (s *Squares) FindPiece(notation string, side Side) (Piece, position.Position) {
	pieces := s.GetPieces(notation, side)
	if len(pieces) == 0 {
		return nil, position.NewEmpty()
	}

	return pieces[0], s.GetByPiece(pieces[0])
}

func (s *Squares) MovePiece(from, to position.Position) (capturedPiece Piece, err error) {
	if !from.IsBoundaries(s.edgePosition) || !to.IsBoundaries(s.edgePosition) {
		return nil, ErrSquareOutOfRange
	}

	capturedPiece = s.rows[to.Rank-1][to.File-1]
	s.rows[from.Rank-1][from.File-1], s.rows[to.Rank-1][to.File-1] = nil, s.rows[from.Rank-1][from.File-1]

	return
}

func (s *Squares) MovePieceTemporarily(from, to position.Position, callback func()) error {
	if !from.IsBoundaries(s.edgePosition) || !to.IsBoundaries(s.edgePosition) {
		return ErrSquareOutOfRange
	}

	fromSquarePiece := s.rows[from.Rank-1][from.File-1]
	toSquarePiece := s.rows[to.Rank-1][to.File-1]

	defer func() {
		s.rows[from.Rank-1][from.File-1], s.rows[to.Rank-1][to.File-1] = fromSquarePiece, toSquarePiece
	}()

	s.rows[from.Rank-1][from.File-1], s.rows[to.Rank-1][to.File-1] = nil, fromSquarePiece

	callback()

	return nil
}

// PlacePiece устанавливает фигуру на клетку по заданной позиции.
func (s *Squares) PlacePiece(piece Piece, position position.Position) error {
	if !position.IsBoundaries(s.edgePosition) {
		return ErrSquareOutOfRange
	}

	s.rows[position.Rank-1][position.File-1] = piece

	return nil
}
