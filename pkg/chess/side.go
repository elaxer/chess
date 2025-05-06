package chess

// Side представляет сторону доски.
type Side bool

const (
	SideWhite Side = true
	SideBlack Side = false
)

func (s Side) IsWhite() bool {
	return s == SideWhite
}

func (s Side) IsBlack() bool {
	return s == SideBlack
}

func (s Side) String() string {
	if s.IsBlack() {
		return "black"
	}

	return "white"
}
