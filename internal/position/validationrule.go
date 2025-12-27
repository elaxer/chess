package position

import "errors"

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
