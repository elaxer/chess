package position

import (
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
)

const (
	// FileNull has a zero value and represents an empty file.
	// This value is considered valid.
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

	// FileMin is the minimum file value after FileNull.
	FileMin = FileA

	// FileMax is the maximum file value supported by the engine.
	// File values greater than FileMax are considered invalid.
	FileMax = FileP
)

const files = "abcdefghijklmnop"

// File represents the horizontal coordinate on the board.
// It takes values from 1 to 16, where 1 corresponds to FileA and 16 to FileP.
// FileNull is a special value representing an uninitialized file, but is still considered valid.
//
// A file may be valid or invalid.
type File int8

// FileFromString returns a File from the specified string.
// The string must be a single character from "a" to "p" (case-insensitive).
// If the input is invalid, the function returns FileNull.
func FileFromString(str string) File {
	if len(str) != 1 {
		return FileNull
	}

	idx := strings.Index(files, strings.ToLower(str))

	return File(idx + 1)
}

// IsNull reports whether the file is FileNull.
func (f File) IsNull() bool {
	return f == FileNull
}

// Validate checks whether the file value is within the range from FileNull to FileMax.
// Returns an error if the value is invalid; otherwise returns nil.
// FileNull is considered valid.
func (f File) Validate() error {
	return validation.Errors{
		"file": validation.Validate(int8(f), validation.Min(int8(FileNull)), validation.Max(int8(FileMax))),
	}.Filter()
}

// String returns the string representation of the file.
// If the file is null or invalid, it returns an empty string.
// Otherwise, it returns the alphabetic representation, e.g., "a" for FileA, "b" for FileB, and so on.
func (f File) String() string {
	if f.IsNull() || f.Validate() != nil {
		return ""
	}

	return string(files[f-1])
}
