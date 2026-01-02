package position

import (
	"errors"
	"strings"
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

const files = "abcdefghijklmnop"

var errFileInvalid = errors.New("the file is invalid")

type File int8

// FileFromString returns a File from the specified string.
// The string must be a single character from "a" to "p" (case-insensitive).
// If the input is invalid, the function returns FileNull.
func FileFromString(str string) File {
	if len(str) != 1 {
		return FileNull
	}

	idx := strings.Index(files, strings.ToLower(str))
	if idx == -1 {
		return FileNull
	}

	//nolint:gosec
	return File(min(idx+1, int(FileMax)))
}

// IsNull reports whether the file is FileNull.
func (f File) IsNull() bool {
	return f == FileNull
}

// Validate checks whether the file value is within the range from FileNull to FileMax.
// Returns an error if the value is invalid; otherwise returns nil.
// FileNull is considered valid.
func (f File) Validate() error {
	if f < FileNull || f > FileMax {
		return errFileInvalid
	}

	return nil
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
