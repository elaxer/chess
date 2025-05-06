package move

const (
	notationCheck = "+"
	notationMate  = "#"
)

type CheckMate struct {
	IsCheck bool
	IsMate  bool
}

func NewCheckMate(notation string) *CheckMate {
	return &CheckMate{notation == notationCheck, notation == notationMate}
}

func (m *CheckMate) String() string {
	switch {
	case m.IsMate:
		return notationMate
	case m.IsCheck:
		return notationCheck
	default:
		return ""
	}
}
