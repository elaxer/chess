package move

import (
	"regexp"

	"github.com/elaxer/chess/pkg/rgx"
)

var RegexpCastling = regexp.MustCompile("^[0Oo]-[0Oo](?P<long>-[0Oo])?[#+]?$")

type Castling bool

const (
	CastlingShort Castling = true
	CastlingLong  Castling = false
)

func CastlingFromString(str string) (Castling, error) {
	result, err := rgx.Group(RegexpCastling, str)
	if err != nil {
		return false, err
	}

	return Castling(result["long"] == ""), nil
}

func (m Castling) Validate() error {
	return nil
}

func (m Castling) IsShort() bool {
	return m == CastlingShort
}

func (m Castling) IsLong() bool {
	return m == CastlingLong
}

func (m Castling) String() string {
	return map[Castling]string{
		CastlingShort: "O-O",
		CastlingLong:  "O-O-O",
	}[m]
}
