package chess

import (
	"fmt"

	"github.com/elaxer/chess/pkg/chess/position"
)

// Piece interface describes a chess piece.
// It provides methods to get the piece's side, check if it has been moved,
// get pseudo-legal moves, and retrieve its notation and weight.
// It is used to represent different types of chess pieces such as pawns, knights, bishops, etc.
// Each piece implements this interface to provide its specific behavior and properties.
// The Piece interface is essential for the chess game logic, allowing the game to handle pieces generically
// while still respecting the unique movement and rules associated with each type of piece.
type Piece interface {
	fmt.Stringer
	// Side returns the side of the piece.
	Side() Side
	// IsMoved returns true if the piece has been moved.
	// This is can be used to determine if the piece can perform castling or en passant.
	IsMoved() bool
	// MarkMoved marks the piece as moved.
	MarkMoved()
	// PseudoMoves returns all pseudo-legal moves for the piece from the given position.
	// Pseudo-move is a move that does not check for checks or other game rules,
	// but only considers the piece's movement capabilities.
	// It returns a set of positions that the piece can move to from the given position.
	PseudoMoves(from position.Position, squares *Squares) position.Set
	// Notation returns the algebraic notation of the piece.
	// For example, for a pawn it returns "", for a knight it returns "N", etc.
	Notation() string
	// Weight returns the weight of the piece.
	// The weight is used to evaluate the piece's value in the game.
	Weight() uint8
}
