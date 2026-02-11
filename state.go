package chess

import "encoding/json"

// StateClear represents a clear state of the chess board.
// This state indicates that there are no threats or special conditions on the board.
var StateClear = NewState("clear", false)

// State represents the type of a board state.
// It is used to categorize the state of the chess board.
type State interface {
	// Name returns the name of the state.
	Name() string
	// IsTerminal indicates that the board has reached a terminal state,
	// where no further moves can be made.
	IsTerminal() bool
}

type state struct {
	name       string
	isTerminal bool
}

// NewState is used to create a new state that can be used in the chess board.
// It allows for the creation of custom states which can be useful
// for representing different conditions on the chess board.
func NewState(name string, isTerminal bool) State {
	return &state{name, isTerminal}
}

func (s *state) Name() string {
	return s.name
}

func (s *state) IsTerminal() bool {
	return s.isTerminal
}

func (s *state) String() string {
	return s.name
}

func (s *state) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"name":        s.name,
		"is_terminal": s.isTerminal,
	})
}
