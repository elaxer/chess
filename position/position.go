// Package position provides functionality to work with chess positions.
// It defines the Position type, which represents a square on a chessboard,
// and provides methods to create, validate, and manipulate positions.
package position

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"

	"github.com/elaxer/rgx"
	validation "github.com/go-ozzo/ozzo-validation"

	mapset "github.com/deckarep/golang-set/v2"
)

// Regexp is a regular expression pattern used to parse chess positions.
// It matches a string that represents a position on the chessboard,
// such as "e4", "a1", or "h8". The pattern captures the file (a-h) and rank (1-8 or 10-16).
var Regexp = regexp.MustCompile("^(?P<file>[a-p])?(?P<rank>1[0-6]|[1-9])?$")

// Set is a set of chess positions.
// It is implemented using the mapset package, which provides a set data structure.
// The Set type is used to store unique positions on the chessboard.
type Set = mapset.Set[Position]

// Position represents the coordinates of a square on a chessboard.
// Position consists of File and Rank.
// Positions can have different states:
//   - Full: File and Rank are filled and both values are not null.
//   - Partial, not empty: File or Rank has null value.
//   - Empty: File and Rank both have null value.
//   - Invalid: File or Rank is invalid (see File and Rank documentation).
type Position struct {
	File File `json:"file"`
	Rank Rank `json:"rank"`
}

func New(file File, rank Rank) Position {
	return Position{file, rank}
}

// NewEmpty creates a new empty position.
// File and rank will have null values.
func NewEmpty() Position {
	return Position{}
}

// FromString creates a new position from chess notation.
// If the string is invalid or it's empty, it returns an empty position.
func FromString(str string) Position {
	data, err := rgx.Group(Regexp, str)
	if err != nil {
		return NewEmpty()
	}

	rank, _ := strconv.Atoi(data["rank"])

	return Position{FileFromString(data["file"]), Rank(rank)}
}

// IsFull checks if the position contains both File and Rank not empty values.
func (p Position) IsFull() bool {
	return !p.File.IsNull() && !p.Rank.IsNull()
}

// IsEmpty checks if the position contains both File and Rank null values.
func (p Position) IsEmpty() bool {
	return p.File.IsNull() && p.Rank.IsNull()
}

// Validate checks if the position contains both File and Rank values not exceeding their maximum limits.
func (p Position) Validate() error {
	return validation.ValidateStruct(&p, validation.Field(&p.File), validation.Field(&p.Rank))
}

// String returns the string representation of the position.
// If the position is empty, it returns an empty string.
// For example, Position{FileE, Rank4} will be converted to "e4".
func (p Position) String() string {
	return fmt.Sprintf("%s%s", p.File, p.Rank)
}

func (p *Position) UnmarshalJSON(data []byte) error {
	position := make(map[string]any, 2)
	if err := json.Unmarshal(data, &position); err != nil {
		return err
	}

	if file, ok := position["file"].(float64); ok {
		if File(file).Validate() == nil {
			p.File = File(file)
		}
	}

	if rank, ok := position["rank"].(float64); ok {
		if Rank(rank).Validate() == nil {
			p.Rank = Rank(rank)
		}
	}

	return nil
}
