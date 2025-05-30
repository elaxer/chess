package move

import (
	"fmt"
	"regexp"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/rgx"
	validation "github.com/go-ozzo/ozzo-validation"
)

var regexpNormal = regexp.MustCompile(fmt.Sprintf(
	"^(?P<piece>[KQBNR])?%s?(?P<is_capture>x)?%s%s?$",
	position.RegexpFrom,
	position.RegexpTo,
	RegexpCheckMate,
))

// Normal представляет обычный ход фигурой в шахматах.
type Normal struct {
	CheckMate
	// PieceNotation обозначает фигуру, которая делает ход.
	PieceNotation string
	// From, To означают начальную и конечную позиции хода.
	From, To position.Position
	// IsCapture означает, было ли взятие фигуры противника в результате хода.
	IsCapture     bool
	CapturedPiece chess.Piece
}

// NormalFromNotation создает новый ход из шахматной нотации.
func NormalFromNotation(notation string) (*Normal, error) {
	data, err := rgx.Group(regexpNormal, notation)
	if err != nil {
		return nil, err
	}

	return &Normal{
		CheckMateFromNotation(data["checkmate"]),
		data["piece"],
		position.FromNotation(data["from"]),
		position.FromNotation(data["to"]),
		data["is_capture"] != "",
		nil,
	}, nil
}

func (m *Normal) Notation() string {
	str := string(m.PieceNotation) + m.From.String()
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
