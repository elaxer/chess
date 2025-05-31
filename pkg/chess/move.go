package chess

import "fmt"

// Move это интерфейс, представляющий ход в шахматной игре.
type Move interface {
	fmt.Stringer
	// Validate проверяет, корректны ли данные хода.
	Validate() error
}

// RawMove это тип хода, который содержит нотацию.
// Это полезно использовать для совершения ходов, когда имеется только строка нотации.
// Релазиует интерфейс Move
type RawMove string

func (m RawMove) String() string {
	return string(m)
}

func (m RawMove) Validate() error {
	return nil
}
