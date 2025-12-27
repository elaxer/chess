package chess

import (
	"regexp"
	"strconv"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/elaxer/chess/internal/position"
	"github.com/elaxer/rgx"
)

const (
	// FileNull has a zero value and represents an empty file.
	// This value is considered valid.
	FileNull = position.FileNull
	// FileMin is the minimum file value after FileNull.
	FileMin = position.FileMin
	// FileMax is the maximum file value supported by the engine.
	// File values greater than FileMax are considered invalid.
	FileMax = position.FileMax

	FileA = position.FileA
	FileB = position.FileB
	FileC = position.FileC
	FileD = position.FileD
	FileE = position.FileE
	FileF = position.FileF
	FileG = position.FileG
	FileH = position.FileH
	FileI = position.FileI
	FileJ = position.FileJ
	FileK = position.FileK
	FileL = position.FileL
	FileM = position.FileM
	FileN = position.FileN
	FileO = position.FileO
	FileP = position.FileP

	// RankNull has a zero value and represents an empty rank.
	// This value considered valid.
	RankNull = position.RankNull
	// RankMin is the minimum rank value after RankNull.
	RankMin = position.RankMin
	// RankMax is the maximum rank value supported by the engine.
	// Rank values greater than RankMax are considered invalid.
	RankMax = position.RankMax

	Rank1  = position.Rank1
	Rank2  = position.Rank2
	Rank3  = position.Rank3
	Rank4  = position.Rank4
	Rank5  = position.Rank5
	Rank6  = position.Rank6
	Rank7  = position.Rank7
	Rank8  = position.Rank8
	Rank9  = position.Rank9
	Rank10 = position.Rank10
	Rank11 = position.Rank11
	Rank12 = position.Rank12
	Rank13 = position.Rank13
	Rank14 = position.Rank14
	Rank15 = position.Rank15
	Rank16 = position.Rank16
)

// regexpPosition is a regular expression pattern used to parse chess positions.
// It matches a string that represents a position on the chessboard,
// such as "e4", "a1", or "h8". The pattern captures the file (a-h) and rank (1-8 or 10-16).
var regexpPosition = regexp.MustCompile("^(?P<file>[a-p])?(?P<rank>1[0-6]|[1-9])?$")

// Position represents the coordinates of a square on a chessboard.
// Position consists of File and Rank.
// Positions can have different states:
//   - Full: File and Rank are filled and both values are not null.
//   - Partial, not empty: File or Rank has null value.
//   - Empty: File and Rank both have null value.
//   - Invalid: File or Rank is invalid (see File and Rank documentation).
type Position = position.Position

// File represents the horizontal coordinate on the board.
// It takes values from 1 to 16, where 1 corresponds to FileA and 16 to FileP.
// FileNull is a special value representing an uninitialized file, but is still considered valid.
//
// A file may be valid or invalid.
type File = position.File

// Rank represents the horizontal coordinate on the board.
// It takes values from 1 to 16, where 1 corresponds to Rank1 and 16 to Rank16.
// RankNull is a special value representing an uninitialized rank, but is still considered valid.
//
// A rank may be valid or invalid.
type Rank = position.Rank

// PositionSet is a set of chess positions.
// It is implemented using the mapset package, which provides a set data structure.
// The Set type is used to store unique positions on the chessboard.
type PositionSet = mapset.Set[Position]

func NewPosition(file File, rank Rank) Position {
	return Position{File: file, Rank: rank}
}

// NewEmptyPosition creates a new empty position.
// File and rank will have null values.
func NewEmptyPosition() Position {
	return Position{}
}

// PositionFromString creates a new position from chess notation.
// If the string is invalid or it's empty, it returns an empty position.
func PositionFromString(str string) Position {
	data, err := rgx.Group(regexpPosition, str)
	if err != nil {
		return NewEmptyPosition()
	}

	file := position.FileFromString(data["file"])
	rank := RankNull

	rankInt, _ := strconv.Atoi(data["rank"])
	if rankInt >= int(RankMin) && rankInt <= int(RankMax) {
		rank = Rank(rankInt)
	}

	return NewPosition(file, rank)
}

// ValidationRulePositionIsEmpty checks if the position is empty.
// It implements validation rule for ozzo-validation.
func ValidationRulePositionIsEmpty(pos any) error {
	return position.ValidationRuleIsEmpty(pos)
}
