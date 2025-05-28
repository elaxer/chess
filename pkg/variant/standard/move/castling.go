package move

import (
	"fmt"
	"regexp"

	"github.com/elaxer/chess/pkg/rgx"
)

var RegexpCastling = regexp.MustCompile(fmt.Sprintf("^0-0(?P<long>-0)?%s?$", RegexpCheckMate))

type Castling struct {
	*CheckMate
	CastlingType
}

func NewCastling(notation string) (*Castling, error) {
	result, err := rgx.Group(RegexpCastling, notation)
	if err != nil {
		return nil, err
	}

	return &Castling{
		CheckMate:    CheckMateFromNotation(result["checkmate"]),
		CastlingType: result["long"] == "",
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
