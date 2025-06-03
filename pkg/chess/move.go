package chess

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Move это интерфейс, представляющий ход в шахматной игре.
type Move interface {
	fmt.Stringer
	validation.Validatable
}

type MoveResult interface {
	Move
	Move() Move
}

// StringMove это тип хода, который содержит нотацию.
// Это полезно использовать для совершения ходов, когда имеется только строка нотации.
// Релазиует интерфейс Move
type StringMove string

func (m StringMove) String() string {
	return string(m)
}

func (m StringMove) Validate() error {
	return nil
}
