package chess

const (
	ColorWhite Color = true
	ColorBlack Color = false
)

// Color is a type that represents the color of a chess piece, board side or turn.
// It can be either white or black.
// The Color type is used to determine the color of a piece and to check if a side or a turn is white or black.
type Color bool

// IsWhite checks if the color is white.
func (c Color) IsWhite() bool {
	return c == ColorWhite
}

// IsBlack checks if the color is black.
func (c Color) IsBlack() bool {
	return c == ColorBlack
}

func (c Color) String() string {
	if c.IsBlack() {
		return "b"
	}

	return "w"
}
