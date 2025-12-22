package chess

const (
	SideWhite Side = true
	SideBlack Side = false
)

// Side is a type that represents the side of a chess piece.
// It can be either white or black.
// The Side type is used to determine the color of a piece and to check if a side or a turn is white or black.
type Side bool

// IsWhite checks if the side is white.
func (s Side) IsWhite() bool {
	return s == SideWhite
}

// IsBlack checks if the side is black.
func (s Side) IsBlack() bool {
	return s == SideBlack
}

func (s Side) String() string {
	if s.IsBlack() {
		return "b"
	}

	return "w"
}
