package move

import (
	"fmt"
	"regexp"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/rgx"
	validation "github.com/go-ozzo/ozzo-validation"
)

// Promotion представляет ход с превращением пешки в другую фигуру.
// В шахматной нотации он записывается как "e8=Q" или "e7=R+".
type Promotion struct {
	*CheckMate
	From, To position.Position
	NewPiece chess.PieceNotation
}

func NewPromotion(notation string) (*Promotion, error) {
	result, err := rgx.Group(regexp.MustCompile(`^(?P<from_file>[a-h])?(?P<to>[a-h][18])=(?P<piece>[QBNR])(?P<checkmate>[+#])?$`), notation)
	if err != nil {
		return nil, err
	}

	return &Promotion{
		CheckMate: NewCheckMate(result["checkmate"]),
		From:      position.FromNotation(result["from_file"]),
		To:        position.FromNotation(result["to"]),
		NewPiece:  chess.PieceNotation(result["piece"]),
	}, nil
}

func (m *Promotion) Notation() string {
	return fmt.Sprintf("%s%s=%s%s", m.From, m.To, m.NewPiece, m.CheckMate)
}

func (m *Promotion) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.To, validation.Required),
		validation.Field(
			&m.NewPiece,
			validation.Required,
			validation.In(chess.NotationQueen, chess.NotationRook, chess.NotationBishop, chess.NotationKnight),
		),
	)
}
