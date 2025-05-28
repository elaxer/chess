package chess

const (
	StateTypeClear StateType = iota
	StateTypeThreat
	StateTypeTerminal
	StateTypeDraw
)

type StateType uint8

func (t StateType) IsClear() bool {
	return t == StateTypeClear
}

func (t StateType) IsThreat() bool {
	return t == StateTypeThreat
}

func (t StateType) IsTerminal() bool {
	return t == StateTypeTerminal || t == StateTypeDraw
}

func (t StateType) IsDraw() bool {
	return t == StateTypeDraw
}
