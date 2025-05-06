package chess

const (
	StateClear State = iota
	StateCheck
	StateMate
	StateStalemate
	StateDraw
)

// State представляет состояние доски.
// Возможные состояния:
// - Clear: обычная позиция
// - Check: шах
// - Mate: мат
// - Stalemate: пат
// - Draw: ничья
type State uint8

func (s State) IsClear() bool {
	return s == StateClear
}

func (s State) IsCheck() bool {
	return s == StateCheck
}

func (s State) IsMate() bool {
	return s == StateMate
}

func (s State) IsStalemate() bool {
	return s == StateStalemate
}

func (s State) IsDraw() bool {
	return s == StateDraw
}
