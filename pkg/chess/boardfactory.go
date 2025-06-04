package chess

// BoardFactory represents a factory for creating chess boards with different ways.
// It provides methods to create an empty board, a filled board, or a board from a list of moves.
// This allows for flexibility in initializing the board state for different scenarios.
type BoardFactory interface {
	// CreateEmpty creates a new empty board with the specified side to move first.
	// It initializes the board without any pieces, ready for a new game.
	CreateEmpty(turn Side) Board
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
