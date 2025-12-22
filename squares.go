package chess

import (
	"errors"
	"fmt"
	"iter"
	"slices"

	"github.com/elaxer/chess/position"
)

// MaxSupportedPosition is the maximum supported position for the chess board.
// It represents the bottom-right corner of the board, which is typically the
// highest rank and file in standard chess notation.
// The engine does not support positions greater than this value.
var MaxSupportedPosition = position.New(position.FileMax, position.RankMax)

// ErrSquareOutOfRange is returned when given square position is out of the valid range.
var ErrSquareOutOfRange = errors.New("square position is out of range")

// Squares represents squares on a chess board.
// It is a 2D slice of Piece, where each Piece represents a chess piece on the board.
// Squares provides methods to iterate over the pieces, find pieces by position,
// get pieces by type and side, move pieces, and place pieces on the squares.
type Squares struct {
	rows         [][]Piece
	edgePosition position.Position
}

// NewSquares creates a new Squares instance with the specified edge position.
// The edge position defines the size of the squares, which is the maximum rank and file.
// If the edge position exceeds the maximum supported position by the engine, it panics.
func NewSquares(edgePosition position.Position) *Squares {
	if edgePosition.File > MaxSupportedPosition.File ||
		edgePosition.Rank > MaxSupportedPosition.Rank {
		panic("cannot create squares with size greater than max size")
	}

	squares := make([][]Piece, edgePosition.Rank)
	for i := range squares {
		squares[i] = make([]Piece, edgePosition.File)
	}

	return &Squares{squares, edgePosition}
}

// SquaresFromPlacement creates a new Squares instance from a given placement map.
// The placement map is a mapping of position to Piece, where each Piece is placed at its corresponding position.
// If the placement is nil or empty, it returns an empty Squares instance with the specified edge position.
// If any piece in the placement is out of boundaries defined by the edge position, it returns an error.
// If edgePosition is not bounded by the maximum supported position, it panics.
func SquaresFromPlacement(
	edgePosition position.Position,
	placement map[position.Position]Piece,
) (*Squares, error) {
	squares := NewSquares(edgePosition)
	if placement == nil {
		return squares, nil
	}

	for position, piece := range placement {
		if err := squares.PlacePiece(piece, position); err != nil {
			return nil, err
		}
	}

	return squares, nil
}

// EdgePosition returns the edge position of the squares.
func (s *Squares) EdgePosition() position.Position {
	return s.edgePosition
}

// Iter iterates over all squares and yields each position and the piece at that position.
// It starts from the top-left corner (FileA, Rank1) and goes to the bottom-right corner (edgePosition).
func (s *Squares) Iter() iter.Seq2[position.Position, Piece] {
	return func(yield func(position position.Position, piece Piece) bool) {
		for rank, row := range s.rows {
			for file, piece := range row {
				//nolint:gosec
				if !yield(position.New(position.File(file+1), position.Rank(rank+1)), piece) {
					return
				}
			}
		}
	}
}

// IterByRows iterates over the squares by rows.
// It yields each row starting from the top (Rank1) to the bottom (edgePosition.Rank).
// If backward is true, it iterates from the bottom to the top.
// Each row is yielded as a sequence of file and their corresponding piece.
func (s *Squares) IterByRows(
	backward bool,
) iter.Seq2[position.Rank, iter.Seq2[position.File, Piece]] {
	rows := slices.Clone(s.rows)
	if backward {
		slices.Reverse(rows)
	}

	return func(yield func(position.Rank, iter.Seq2[position.File, Piece]) bool) {
		for rank, row := range rows {
			//nolint:gosec
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

// IterByDirection iterates over the squares in a specific direction from a given position.
// From is the starting position, and direction is the step to take in each iteration.
// The direction is defined by a position.Position, which indicates the change in file and rank.
// It yields each position and the piece at that position until it goes out of boundaries.
// The boundaries are defined by the edgePosition of the squares.
func (s *Squares) IterByDirection(
	from, direction position.Position,
) iter.Seq2[position.Position, Piece] {
	return func(yield func(position position.Position, piece Piece) bool) {
		position := position.New(from.File+direction.File, from.Rank+direction.Rank)
		for s.isBoundaries(position) {
			if !yield(position, s.rows[position.Rank-1][position.File-1]) {
				return
			}

			position.File += direction.File
			position.Rank += direction.Rank
		}
	}
}

// FindByPosition finds a piece by its position on the squares.
// If the position is out of boundaries, it returns ErrSquareOutOfRange.
// If the position is valid, it returns the piece at that position or nil if no piece is found.
func (s *Squares) FindByPosition(position position.Position) (Piece, error) {
	if !s.isBoundaries(position) {
		return nil, fmt.Errorf("%w (%s)", ErrSquareOutOfRange, position)
	}

	return s.rows[position.Rank-1][position.File-1], nil
}

// GetByPiece returns the position of the occupied square by the given piece.
// If the piece is not found on the squares or it is nil, it returns an empty position.
func (s *Squares) GetByPiece(piece Piece) position.Position {
	for rank, row := range s.rows {
		for file := range row {
			if row[file] == piece {
				//nolint:gosec
				return position.New(position.File(file+1), position.Rank(rank+1))
			}
		}
	}

	return position.NewEmpty()
}

// FindPiece finds the first piece of a given type for a specific side and returns it along with its position.
// If no piece is found, it returns nil and an empty position.
func (s *Squares) FindPiece(notation string, side Side) (Piece, position.Position) {
	pieces := s.GetPieces(notation, side)
	if len(pieces) == 0 {
		return nil, position.NewEmpty()
	}

	return pieces[0], s.GetByPiece(pieces[0])
}

// GetPieces returns all pieces of a specific type (notation) for a given side.
// It iterates through all squares and collects pieces that match the given notation and side.
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

// GetAllPieces returns all pieces on the squares for a given side.
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

// PlacePiece places a piece on the squares at the specified position.
// If the position is out of boundaries, it returns ErrSquareOutOfRange.
// The piece is placed at the given position, and any existing piece at that position is overwritten.
// If the piece is nil, it will overwrite the existing piece at that position with nil.
func (s *Squares) PlacePiece(piece Piece, position position.Position) error {
	if !s.isBoundaries(position) {
		return ErrSquareOutOfRange
	}

	s.rows[position.Rank-1][position.File-1] = piece

	return nil
}

// MovePiece moves a piece from one position to another.
// It returns the captured piece if any, or nil if no piece was captured.
// If the move is out of boundaries, it returns ErrSquareOutOfRange.
// The piece at the 'from' position is moved to the 'to' position, and the 'from' position is set to nil.
// If the 'to' position already has a piece, it is captured and returned.
// If the 'from' position is empty, it returns nil and no error.
func (s *Squares) MovePiece(from, to position.Position) (capturedPiece Piece, err error) {
	if !s.isBoundaries(from) || !s.isBoundaries(to) {
		return nil, ErrSquareOutOfRange
	}

	if s.rows[from.Rank-1][from.File-1] == nil {
		return nil, nil
	}

	capturedPiece = s.rows[to.Rank-1][to.File-1]
	s.rows[from.Rank-1][from.File-1], s.rows[to.Rank-1][to.File-1] = nil, s.rows[from.Rank-1][from.File-1]

	return
}

// MovePieceTemporarily moves a piece from one position to another temporarily.
// It allows executing a callback function while the piece is moved.
// After the callback is executed, the piece is moved back to its original position.
// If the move is out of boundaries, it returns ErrSquareOutOfRange.
// This is useful for testing or simulating moves without permanently changing the squares position.
func (s *Squares) MovePieceTemporarily(from, to position.Position, callback func()) error {
	if !s.isBoundaries(from) || !s.isBoundaries(to) {
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

func (s *Squares) isBoundaries(pos position.Position) bool {
	return pos.File >= position.FileMin && pos.File <= s.edgePosition.File &&
		pos.Rank >= position.RankMin && pos.Rank <= s.edgePosition.Rank
}
