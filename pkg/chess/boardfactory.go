package chess

// BoardFactory это интерфейс для создания шахматной доски.
type BoardFactory interface {
	// CreateEmpty создает пустую шахматную доску.
	// Должен возвращать пустую доску с пустыми клетками.
	CreateEmpty(turn Side) Board
	// CreateFilled создает шахматную доску заполненную фигурами.
	CreateFilled() Board
	// CreateFromMoves создает шахматную доску из списка ходов.
	// Должен возвращать доску с фигурами, которые были на ней в момент последнего хода.
	// Если один из ходов окажется невалидным, то метод вернет ошибку.
	CreateFromMoves(moves []Move) (Board, error)
}
