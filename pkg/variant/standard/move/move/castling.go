package move

import (
	"fmt"
	"regexp"

	"github.com/elaxer/chess/pkg/rgx"
)

var regexpCastling = regexp.MustCompile(fmt.Sprintf("^[0Oo]-[0Oo](?P<long>-[0Oo])?%s?$", RegexpSuffix))

type Castling struct {
	abstract
	CastlingType
}

func NewCastling(castlingType CastlingType) *Castling {
	return &Castling{CastlingType: castlingType}
}

func CastlingFromNotation(notation string) (*Castling, error) {
	result, err := rgx.Group(regexpCastling, notation)
	if err != nil {
		return nil, err
	}

	return &Castling{
		abstract:     abstractFromNotation(result["suffix"]),
		CastlingType: result["long"] == "",
	}, nil
}

func (m *Castling) Validate() error {
	return nil
}

func (m *Castling) String() string {
	return m.CastlingType.String() + m.abstract.String()
}
