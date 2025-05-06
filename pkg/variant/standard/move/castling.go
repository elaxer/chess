package move

import (
	"regexp"

	"github.com/elaxer/chess/pkg/rgx"
)

type Castling struct {
	*CheckMate
	CastlingType
}

func NewCastling(notation string) (*Castling, error) {
	result, err := rgx.Group(regexp.MustCompile("^0-0(?P<long_castling>-0)?(?P<checkmate>[+#])?$"), notation)
	if err != nil {
		return nil, err
	}

	return &Castling{
		CheckMate:    NewCheckMate(result["checkmate"]),
		CastlingType: result["long_castling"] == "",
	}, nil
}

func (m *Castling) Notation() string {
	return m.CastlingType.String() + m.CheckMate.String()
}

func (m *Castling) Validate() error {
	return nil
}

func (m *Castling) String() string {
	return m.Notation()
}
