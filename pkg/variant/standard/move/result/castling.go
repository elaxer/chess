package result

import (
	"fmt"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/variant/standard/move/move"
)

type Castling struct {
	Abstract
	move.Castling `json:"castling"`
}

func (r *Castling) Move() chess.Move {
	return r.Castling
}

func (r *Castling) String() string {
	return fmt.Sprintf("%s%s", r.Castling, r.suffix())
}
