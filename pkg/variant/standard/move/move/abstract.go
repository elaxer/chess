package move

import (
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/variant/standard/state/state"
)

const (
	notationCheck = "+"
	notationMate  = "#"
)

const RegexpCheckMate = `(?P<checkmate>[+#])`

type abstract struct {
	NewBoardState chess.State
}

func abstractFromNotation(notation string) abstract {
	switch notation {
	case notationCheck:
		return abstract{state.Check}
	case notationMate:
		return abstract{state.Mate}
	}

	return abstract{nil}
}

func (m abstract) String() string {
	switch m.NewBoardState {
	case state.Mate:
		return notationMate
	case state.Check:
		return notationCheck
	}

	return ""
}
