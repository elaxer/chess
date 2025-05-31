package move

import (
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/variant/standard/state/state"
)

const (
	notationCheck     = "+"
	notationCheckmate = "#"
)

const RegexpSuffix = `(?P<suffix>[+#])`

type abstract struct {
	NewBoardState chess.State
}

func abstractFromNotation(notation string) abstract {
	switch notation {
	case notationCheck:
		return abstract{state.Check}
	case notationCheckmate:
		return abstract{state.Checkmate}
	}

	return abstract{nil}
}

func (m abstract) String() string {
	switch m.NewBoardState {
	case state.Checkmate:
		return notationCheckmate
	case state.Check:
		return notationCheck
	}

	return ""
}
