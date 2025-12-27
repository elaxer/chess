package chess

const (
	// StateTypeClear indicates that the chess board is in a clear state,
	// meaning there are no threats or special conditions affecting the game.
	// This is the default state of the board when no pieces are threatening each other.
	StateTypeClear StateType = iota
	// StateTypeThreat indicates that there is a threat on the chess board,
	// which is useful for indicating check or other conditions where a piece is under threat.
	StateTypeThreat
	// StateTypeTerminal indicates that the game has reached a terminal state,
	// such as checkmate or stalemate, where no further moves can be made.
	// This state is used to signify the end of the game.
	StateTypeTerminal
)

// StateType represents the type of a chess board state.
// It is used to categorize the state of the chess board.
// The StateType can be one of the following:
// - StateTypeClear
// - StateTypeThreat
// - StateTypeTerminal.
type StateType uint8

// IsClear checks if the state type is clear.
func (t StateType) IsClear() bool {
	return t == StateTypeClear
}

// IsThreat checks if the state type indicates a threat on the board.
func (t StateType) IsThreat() bool {
	return t == StateTypeThreat
}

// IsTerminal checks if the state type indicates a terminal state.
func (t StateType) IsTerminal() bool {
	return t == StateTypeTerminal
}
