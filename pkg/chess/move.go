package chess

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Move represents a chess input move which used to make a move on the chessboard.
// It is an interface that can be implemented by different types of moves,
// such as Normal move or Castling and so on.
// The Move interface requires a String method for string representation
// and a Validate method for validation purposes.
type Move interface {
	fmt.Stringer
	validation.Validatable
}

// MoveResult is an interface that represents the result of a move made on the chessboard.
// It includes additional methods for retrieving the move itself and validating the result.
// It is used to encapsulate the details of a move, including the side that made the move,
// the new state of the board after the move, and can includes any additional information such as captured pieces or castling details.
// The MoveResult interface requires a String method for string representation,
// a Validate method for validation purposes, and a Move method to retrieve the move that was made.
type MoveResult interface {
	fmt.Stringer
	validation.Validatable
	// Move returns the Move that was made.
	// This method is used to retrieve the move that resulted in this MoveResult.
	Move() Move
	// Side returns the Side that made the move.
	// This method is used to determine which side (white or black) made the move.
	Side() Side
	// BoardNewState returns the new state of the board after the move.
	BoardNewState() State
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
