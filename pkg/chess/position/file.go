package position

import (
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
)

const (
	FileNull File = iota

	FileA
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH

	FileI
	FileJ
	FileK
	FileL
	FileM
	FileN
	FileO
	FileP

	FileMin = FileA
	FileMax = FileP
)

const (
	RegexpFile = "[a-p]"

	files = "abcdefghijklmnop"
)

// File представляет вертикаль на шахматной доске.
// Принимает значения от 1 до 16, где 1 соответствует вертикали "a", а 16 - вертикали "p".
type File int8

// NewFile создает новый объект File из символа, представляющего вертикаль.
func NewFile(char string) File {
	if char == "" {
		return File(0)
	}

	idx := strings.Index(files, strings.ToLower(char))

	return File(idx + 1)
}

func (f File) IsNull() bool {
	return f == FileNull
}

func (f File) Validate() error {
	return validation.Errors{
		"file": validation.Validate(int8(f), validation.Required, validation.Min(int8(FileMin)), validation.Max(int8(FileMax))),
	}.Filter()
}

func (f File) String() string {
	if f.Validate() != nil {
		return ""
	}

	return string(files[f-1])
}
