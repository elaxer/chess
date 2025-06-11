package chess

import "fmt"

var (
	// StateClear represents a clear state of the chess board.
	// This state indicates that there are no threats or special conditions on the board.
	StateClear = NewState("clear", StateTypeClear)
)

// StateType represents the type of a board state.
// It is used to categorize the state of the chess board.
type State interface {
	fmt.Stringer
	// Type returns the type of the state.
	// The type can be one of the predefined StateType values,
	// such as StateTypeClear, StateTypeThreat, StateTypeTerminal, or StateTypeDraw.
	Type() StateType
}

type state struct {
	name      string
	stateType StateType
}

// NewState creates a new State with the given name and type.
// The name is a string representation of the state,
// and the stateType is one of the predefined StateType values.
// This function is used to create a new state that can be used in the chess game.
// It allows for the creation of custom states with specific names and types,
// which can be useful for representing different conditions on the chess board.
func NewState(name string, stateType StateType) State {
	return &state{
		name:      name,
		stateType: stateType,
	}
}

// Type returns the type of the state.
func (s *state) Type() StateType {
	return s.stateType
}

// String returns the name of the state.
func (s *state) String() string {
	return s.name
}
