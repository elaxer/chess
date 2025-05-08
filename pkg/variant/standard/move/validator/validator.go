// validator содержит валидаторы для проверки возможности выполнения ходов в шахматах.
package validator

import (
	"errors"
	"fmt"
)

var (
	Err                = errors.New("ошибка валидации хода")
	ErrEmptySquare     = fmt.Errorf("%w: клетка не имеет фигуры", Err)
	ErrInvalidNotation = fmt.Errorf("%w: некорректная нотация хода", Err)
)
