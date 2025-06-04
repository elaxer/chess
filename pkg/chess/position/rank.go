package position

import (
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
)

const (
	// RankNull has a zero value and represents an empty rank.
	// This value considered valid.
	RankNull Rank = iota

	Rank1
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

	// RankMin is the minimum rank value after RankNull.
	RankMin = Rank1

	// RankMax is the maximum rank value supported by the engine.
	// Rank values greater than RankMax are considered invalid.
	RankMax = Rank16
)

// Rank represents the horizontal coordinate on the board.
// It takes values from 1 to 16, where 1 corresponds to Rank1 and 16 to Rank16.
// RankNull is a special value representing an uninitialized rank, but is still considered valid.
//
// A rank may be valid or invalid.
type Rank int8

// IsNull reports whether the rank is RankNull.
func (r Rank) IsNull() bool {
	return r == RankNull
}

// Validate checks whether the rank value exceeds RankMax.
// Returns an error if the value is invalid; otherwise returns nil.
// RankNull is considered valid.
func (r Rank) Validate() error {
	return validation.Errors{
		"rank": validation.Validate(int8(r), validation.Max(int8(RankMax))),
	}.Filter()
}

// String returns the string representation of the rank.
// If the rank is null or invalid, it returns an empty string.
// Otherwise, it returns the numeric representation, e.g., "1" for Rank1, "2" for Rank2, and so on.
func (r Rank) String() string {
	if r.IsNull() || r.Validate() != nil {
		return ""
	}

	return strconv.Itoa(int(r))
}
