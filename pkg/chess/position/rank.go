package position

import (
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
)

const (
	Rank1 Rank = iota + 1
	Rank2
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8

	Rank9
	Rank10
	Rank11
	Rank12
	Rank13
	Rank14
	Rank15
	Rank16

	RankMin = Rank1
	RankMax = Rank16
)

const RegexpRank = "(1[0-6]|[1-9])"

// Rank представляет горизонталь на шахматной доске.
// Принимает значения от 1 до 16.
type Rank int8

func (r Rank) Validate() error {
	return validation.Errors{
		"rank": validation.Validate(int8(r), validation.Required, validation.Min(int8(RankMin)), validation.Max(int8(RankMax))),
	}.Filter()
}

func (r Rank) String() string {
	if r.Validate() != nil {
		return ""
	}

	return strconv.Itoa(int(r))
}
