package chess

import "fmt"

// Move represents a chess input move which used to make a move on the chessboard.
// It is an interface that can be implemented by different types of moves,
// such as Normal move or Castling and so on.
// The Move interface requires a String method for string representation
// and a Validate method for validation purposes.
type Move interface {
	fmt.Stringer
}

// StringMove is a simple implementation of the Move interface
// that represents a move as a string.
// It is used for cases where the move can be represented as a string,
// such as in standard chess notation.
type StringMove string

func (m StringMove) String() string {
	return string(m)
}

// Validate implements the validation.Validatable interface for StringMove.
// It does not perform any validation and always returns nil.
func (m StringMove) Validate() error {
	return nil
}
