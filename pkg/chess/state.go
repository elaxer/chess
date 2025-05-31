package chess

import "fmt"

var (
	StateClear = NewState("clear", StateTypeClear)
)

type State interface {
	fmt.Stringer
	Type() StateType
}

type state struct {
	name      string
	stateType StateType
}

func NewState(name string, stateType StateType) State {
	return &state{
		name:      name,
		stateType: stateType,
	}
}

func (s *state) Type() StateType {
	return s.stateType
}

func (s *state) String() string {
	return s.name
}
