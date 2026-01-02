package position

import (
	"errors"
	"strconv"
)

const (
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

	RankMin = Rank1

	RankMax = Rank16
)

var errRankInvalid = errors.New("the rank is invalid")

type Rank int8

// IsNull reports whether the rank is RankNull.
func (r Rank) IsNull() bool {
	return r == RankNull
}

// Validate checks whether the rank value is within the range from RankNull to RankMax.
// Returns an error if the value is invalid; otherwise returns nil.
// RankNull is considered valid.
func (r Rank) Validate() error {
	if r < RankNull || r > RankMax {
		return errRankInvalid
	}

	return nil
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
