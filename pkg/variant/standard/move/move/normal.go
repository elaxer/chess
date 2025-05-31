package move

import (
	"fmt"
	"regexp"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/rgx"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
	validation "github.com/go-ozzo/ozzo-validation"
)

var regexpNormal = regexp.MustCompile(fmt.Sprintf(
	"^(?P<piece>[KQBNR])?%s?(?P<is_capture>x)?%s%s?$",
	position.RegexpFrom,
	position.RegexpTo,
	RegexpSuffix,
))

// Normal представляет обычный ход фигурой в шахматах.
type Normal struct {
	abstract
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
		abstractFromNotation(data["suffix"]),
		data["piece"],
		position.FromString(data["from"]),
		position.FromString(data["to"]),
		data["is_capture"] != "",
		nil,
	}, nil
}

func (m *Normal) Validate() error {
	pieceNotations := make([]any, 0, len(piece.AllNotations))
	for _, notation := range piece.AllNotations {
		pieceNotations = append(pieceNotations, notation)
	}

	return validation.ValidateStruct(
		m,
		validation.Field(&m.PieceNotation, validation.In(pieceNotations...)),
		validation.Field(&m.From),
		validation.Field(&m.To),
	)
}

func (m *Normal) String() string {
	capture := ""
	if m.IsCapture {
		capture = "x"
	}

	return fmt.Sprintf("%s%s%s%s%s", m.PieceNotation, m.From, capture, m.To, m.abstract)
}
