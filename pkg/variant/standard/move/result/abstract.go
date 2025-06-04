package result

import (
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/variant/standard/state/state"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Abstract struct {
	MoveSide chess.Side  `json:"side"`
	NewState chess.State `json:"board_new_state"`
}

func (r Abstract) Side() chess.Side {
	return r.MoveSide
}

func (r Abstract) BoardNewState() chess.State {
	return r.NewState
}

func (r Abstract) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.NewState, validation.Required),
	)
}

func (r Abstract) suffix() string {
	switch r.NewState {
	case state.Check:
		return "+"
	case state.Checkmate:
		return "#"
	}

	return ""
}
