package position

import (
	"errors"
)

func RuleIsNotNull(position any) error {
	pos, ok := position.(Position)
	if !ok {
		return nil
	}

	if pos.IsEmpty() {
		return errors.New("position cannot be null")
	}

	return nil
}
