package result

import (
	"fmt"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/variant/standard/move/move"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Promotion struct {
	Piece
	InputMove move.Promotion `json:"input"`
}

func (r *Promotion) Move() chess.Move {
	return &r.InputMove
}

func (r *Promotion) Validate() error {
	return validation.ValidateStruct(
		r,
		validation.Field(&r.Piece),
		validation.Field(&r.InputMove),
	)
}

func (r *Promotion) String() string {
	from := r.FromShortened
	if from.IsEmpty() && r.IsCapture() {
		from.File = r.FromFull.File
	}

	return fmt.Sprintf("%s%s%s=%s%s", from, r.captureString(), r.InputMove.To, r.InputMove.PromotedPieceNotation, r.suffix())
}
