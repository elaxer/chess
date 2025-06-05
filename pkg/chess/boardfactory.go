package chess

import "github.com/elaxer/chess/pkg/chess/position"

// BoardFactory represents a factory for creating chess boards with different ways.
// It provides methods to create an empty board, a filled board, or a board from a list of moves.
// This allows for flexibility in initializing the board state for different scenarios.
type BoardFactory interface {
	// Create creates a new empty board with the specified side to move first and a placement map.
	// If the placement map is nil, it initializes an empty board.
	// If some positions in the placement map are not valid, it returns an error.
	Create(turn Side, placement map[position.Position]Piece) (Board, error)
	// CreateFilled creates a new filled board with the standard chess setup.
	// It initializes the board with all pieces in their starting positions.
	// It does not take any parameters and returns a fully set up board.
	// This is typically used to start a new game of chess.
	CreateFilled() Board
	// CreateFromMoves creates a new board from a list of moves.
	// It initializes the board based on the moves provided, allowing for a board position
	// that reflects a specific sequence of moves.
	// It takes a slice of Move as input and returns a Board that reflects the position after those moves.
	// If the moves are invalid or cannot be applied, it returns an error.
	CreateFromMoves(moves []Move) (Board, error)
}
