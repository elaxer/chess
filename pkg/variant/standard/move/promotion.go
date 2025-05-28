package move

import (
	"fmt"
	"regexp"

	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/rgx"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
	validation "github.com/go-ozzo/ozzo-validation"
)

var promotionRegexp = regexp.MustCompile(fmt.Sprintf(
	`^(?P<from_file>[a-p])?(?P<is_capture>x)?%s=(?P<piece>[QBNR])%s?$`,
	position.RegexpTo,
	RegexpCheckMate,
))

// Promotion представляет ход с превращением пешки в другую фигуру.
// В шахматной нотации он записывается как "e8=Q" или "e7=R+".
type Promotion struct {
	*Normal
	NewPieceNotation string
}

func NewPromotion(notation string) (*Promotion, error) {
	result, err := rgx.Group(promotionRegexp, notation)
	if err != nil {
		return nil, err
	}

	return &Promotion{
		&Normal{
			CheckMateFromNotation(result["checkmate"]),
			"",
			position.FromNotation(result["from_file"]),
			position.FromNotation(result["to"]),
			result["is_capture"] != "",
			nil,
		},
		result["piece"],
	}, nil
}

func (m *Promotion) Notation() string {
	return fmt.Sprintf("%s%s=%s%s", m.From, m.To, m.NewPieceNotation, m.CheckMate)
}

func (m *Promotion) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.To, validation.Required),
		validation.Field(
			&m.NewPieceNotation,
			validation.Required,
			validation.In(piece.NotationQueen, piece.NotationRook, piece.NotationBishop, piece.NotationKnight),
		),
	)
}
