package result

import (
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/variant/standard/state/state"
)

type Abstract struct {
	Side          chess.Side  `json:"side"`
	BoardNewState chess.State `json:"board_new_state"`
}

func (r Abstract) suffix() string {
	switch r.BoardNewState {
	case state.Check:
		return "+"
	case state.Checkmate:
		return "#"
	}

	return ""
}
