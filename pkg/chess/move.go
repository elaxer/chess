package chess

// Move это интерфейс, представляющий ход в шахматной игре.
type Move interface {
	// Notation возвращает нотацию хода.
	Notation() string
	// Validate проверяет, корректны ли данные хода.
	Validate() error
}

// RawMove это тип хода, который содержит нотацию.
// Это полезно использовать для совершения ходов, когда имеется только строка нотации.
// Релазиует интерфейс Move
type RawMove string

func (m RawMove) Notation() string {
	return string(m)
}

func (m RawMove) Validate() error {
	return nil
}
