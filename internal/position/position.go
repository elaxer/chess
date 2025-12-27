// Package position provides functionality to work with chess positions.
// It defines the Position type, which represents a square on a chessboard,
// and provides methods to create, validate, and manipulate positions.
package position

import (
	"encoding/json"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Position represents the coordinates of a square on a chessboard.
//
//nolint:recvcheck
type Position struct {
	File File `json:"file"`
	Rank Rank `json:"rank"`
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
