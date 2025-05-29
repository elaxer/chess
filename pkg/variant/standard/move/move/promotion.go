package move

import (
	"fmt"
	"regexp"

	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/rgx"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
	validation "github.com/go-ozzo/ozzo-validation"
)

var regexpPromotion = regexp.MustCompile(fmt.Sprintf(
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

func PromotionFromNotation(notation string) (*Promotion, error) {
	data, err := rgx.Group(regexpPromotion, notation)
	if err != nil {
		return nil, err
	}

	return &Promotion{
		&Normal{
			CheckMateFromNotation(data["checkmate"]),
			"",
			position.FromNotation(data["from_file"]),
			position.FromNotation(data["to"]),
			data["is_capture"] != "",
			nil,
		},
		data["piece"],
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

func (m *Promotion) String() string {
	return m.Notation()
}
