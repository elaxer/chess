package move

import (
	"fmt"
	"regexp"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/rgx"
	validation "github.com/go-ozzo/ozzo-validation"
)

var normalRegexp = regexp.MustCompile(fmt.Sprintf(
	"^(?P<piece>[KQBNR])?%s?(?P<is_capture>x)?%s%s?$",
	position.RegexpFrom,
	position.RegexpTo,
	RegexpCheckMate,
))

// Normal представляет обычный ход фигурой в шахматах.
type Normal struct {
	*CheckMate
	// PieceNotation обозначает фигуру, которая делает ход.
	PieceNotation chess.PieceNotation
	// From, To означают начальную и конечную позиции хода.
	From, To position.Position
	// IsCapture означает, было ли взятие фигуры противника в результате хода.
	IsCapture     bool
	CapturedPiece chess.Piece
}

// NewNormal создает новый ход из шахматной нотации.
func NewNormal(notation string) (*Normal, error) {
	result, err := rgx.Group(normalRegexp, notation)
	if err != nil {
		return nil, err
	}

	return &Normal{
		NewCheckMate(result["checkmate"]),
		chess.PieceNotation(result["piece"]),
		position.FromNotation(result["from"]),
		position.FromNotation(result["to"]),
		result["is_capture"] != "",
		nil,
	}, nil
}

func (m *Normal) Notation() string {
	str := string(m.PieceNotation)
	if m.From.Validate() == nil {
		str += m.From.String()
	}
	if m.IsCapture {
		str += "x"
	}

	return str + m.To.String() + m.CheckMate.String()
}

func (m *Normal) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.To),
	)
}

func (m *Normal) String() string {
	return m.Notation()
}
