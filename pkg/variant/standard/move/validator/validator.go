// validator содержит валидаторы для проверки возможности выполнения ходов в шахматах.
package validator

import (
	"errors"
	"fmt"
)

var (
	Err                = errors.New("validation error")
	ErrEmptySquare     = fmt.Errorf("%w: no piece at square", Err)
	ErrInvalidNotation = fmt.Errorf("%w: invalid move notation", Err)
)
