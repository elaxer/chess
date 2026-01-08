package chess

import (
	"strconv"
	"unicode"

	"github.com/elaxer/chess/internal/position"
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

var (
	DirectionTop    = NewPosition(FileNull, Rank(1))
	DirectionBottom = NewPosition(FileNull, Rank(-1))
	DirectionLeft   = NewPosition(File(-1), RankNull)
	DirectionRight  = NewPosition(File(1), RankNull)

	DirectionTopLeft     = NewPosition(File(-1), Rank(1))
	DirectionTopRight    = NewPosition(File(1), Rank(1))
	DirectionBottomLeft  = NewPosition(File(-1), Rank(-1))
	DirectionBottomRight = NewPosition(File(1), Rank(-1))
)

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

func NewPosition(file File, rank Rank) Position {
	return Position{File: file, Rank: rank}
}

// NewPositionEmpty creates a new empty position.
// File and rank will have null values.
func NewPositionEmpty() Position {
	return Position{}
}

// PositionFromString creates a new position from chess notation.
// If the string is invalid or it's empty, it returns an empty position.
func PositionFromString(str string) Position {
	if len(str) == 0 || len(str) > 3 {
		return NewPositionEmpty()
	}

	file := FileNull
	rank := RankNull

	rankShift := 0

	if unicode.IsLetter(rune(str[0])) {
		file = position.FileFromString(string(str[0]))
		rankShift = 1

		if err := file.Validate(); err != nil {
			return NewPositionEmpty()
		}
	}

	if len(str) == 0+rankShift {
		return NewPosition(file, rank)
	}

	rankStr := ""
	rankStr += string(str[0+rankShift])

	if len(str) == 2+rankShift {
		rankStr += string(str[1+rankShift])
	}

	rankInt, err := strconv.Atoi(rankStr)
	if err != nil {
		return NewPositionEmpty()
	}

	//nolint:gosec
	rank = Rank(rankInt)
	if err := rank.Validate(); err != nil {
		return NewPositionEmpty()
	}

	return NewPosition(file, rank)
}
