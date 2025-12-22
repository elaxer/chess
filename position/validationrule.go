package position

import (
	"errors"
)

// ValidationRuleIsEmpty checks if the position is empty.
// It implements validation rule for ozzo-validation.
func ValidationRuleIsEmpty(position any) error {
	pos, ok := position.(Position)
	if !ok {
		return nil
	}

	if pos.IsEmpty() {
		//nolint:err113
		return errors.New("position cannot be null")
	}

	return nil
}
