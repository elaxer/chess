// Package chess provides interfaces and types for chess game mechanics development.
// It includes definitions for the chess board, pieces, moves, and game state.
// It allows for the creation and manipulation of chess boards, including making moves,
// checking legal moves, and managing the game state.
// The package is designed to be flexible and extensible, supporting various chess variants and rules.
// It is intended for use in chess applications, engines, and libraries that require chess logic.
package chess

import "github.com/elaxer/chess/pkg/chess/position"

// Board represents a chessboard with its squares, pieces, and board state.
// It provides methods to access the current turn, state of the board, and to make moves.
type Board interface {
	// Squares returns a reference to the squares on the board.
	// It allows access to the individual squares and their pieces.
	Squares() *Squares
	// Turn returns the current turn of the board.
	// It indicates which side (white or black) is to move next.
	Turn() Side
	// State returns the current state of the board for the specified side.
	// Returns chess.StateClear if the board is in a clear state.
	// Should not return nil.
	State(side Side) State
	// MovesHistory returns the history of moves made on the board.
	// It returns a slice of MoveResult, which contains the details of each move.
	MovesHistory() []MoveResult
	// Moves returns a set of legal moves for the specified side.
	Moves(side Side) position.Set
	// LegalMoves returns a set of legal moves for the specified piece.
	// If the piece is nil or not found, it returns an empty set.
	LegalMoves(piece Piece) position.Set
	// MakeMove applies a move to the board and returns the result of the move.
	// It returns a MoveResult which contains the details of the move made.
	// If the move is invalid or cannot be made, it returns an error.
	MakeMove(move Move) (MoveResult, error)
	// UndoLastMove undoes the last move made on the board and returns it.
	// It returns an error if the undo operation fails.
	// Returns nil, nil if there are no moves to undo.
	UndoLastMove() (MoveResult, error)
}
