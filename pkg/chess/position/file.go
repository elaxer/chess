package position

import (
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
)

// todo уйти от фиксированных величин
const (
	FileA File = iota + 1
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH
)

const files = "abcdefgh"

// File представляет вертикаль на шахматной доске.
// Принимает значения от 1 до 8, где 1 соответствует вертикали "a", а 8 - вертикали "h".
type File int8

// NewFile создает новый объект File из символа, представляющего вертикаль.
func NewFile(char string) File {
	idx := strings.Index(files, strings.ToLower(char))

	return File(idx + 1)
}

func (f File) Validate() error {
	return validation.Errors{
		"file": validation.Validate(int8(f), validation.Required, validation.Min(1), validation.Max(8)),
	}.Filter()
}

func (f File) String() string {
	if err := f.Validate(); err != nil {
		return "?"
	}

	return string(files[f-1])
}
