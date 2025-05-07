package move

import (
	"regexp"
	"strconv"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/rgx"
	validation "github.com/go-ozzo/ozzo-validation"
)

// Normal представляет обычный ход фигурой в шахматах.
type Normal struct {
	*CheckMate
	// From, To означают начальную и конечную позиции хода.
	From, To position.Position
	// PieceNotation обозначает фигуру, которая делает ход.
	PieceNotation chess.PieceNotation
	// IsCapture означает, было ли взятие фигуры противника в результате хода.
	IsCapture     bool
	CapturedPiece chess.Piece
}

// NewNormal создает новый ход из шахматной нотации.
func NewNormal(notation string) (*Normal, error) {
	const regexpStr = "^(?P<piece>[KQBNR])?(?P<file_from>[a-h])?(?P<rank_from>[1-8])?(?P<is_capture>x)?(?P<to>[a-h][1-8])(?P<checkmate>[+#])?$"
	result, err := rgx.Group(regexp.MustCompile(regexpStr), notation)
	if err != nil {
		return nil, err
	}

	move := &Normal{
		CheckMate:     NewCheckMate(result["checkmate"]),
		To:            position.FromNotation(result["to"]),
		PieceNotation: chess.PieceNotation(result["piece"]),
		IsCapture:     result["is_capture"] != "",
	}

	if result["file_from"] != "" {
		move.From.File = position.NewFile(result["file_from"])
	}

	if result["rank_from"] != "" {
		rankFrom, _ := strconv.Atoi(result["rank_from"])
		move.From.Rank = position.Rank(rankFrom)

	}

	return move, nil
}

func (m *Normal) Notation() string {
	str := string(m.PieceNotation)
	if m.From.File != 0 {
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
