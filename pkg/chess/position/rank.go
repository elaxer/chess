package position

import validation "github.com/go-ozzo/ozzo-validation"

const (
	Rank1 Rank = iota + 1
	Rank2
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8
)

// Rank представляет горизонталь на шахматной доске.
// Принимает значения от 1 до 8.
type Rank int8

func (r Rank) Validate() error {
	return validation.Errors{
		"rank": validation.Validate(int8(r), validation.Required, validation.Min(1), validation.Max(8)),
	}.Filter()
}
