package chess

// PieceFactory is an interface for creating chess pieces from notation or string representations.
// It allows for the creation of pieces based on their notation.
// It also supports creating pieces from a string representation, which may include the piece's side (white or black).
// Implementations of this interface should handle the conversion
// and validation of the input strings to produce the correct Piece instances.
type PieceFactory interface {
	// Create creates a Piece from a given notation string and side.
	// The notation string should be one of the predefined notations.
	// It returns the created Piece or an error if the notation is invalid.
	Create(notation string, side Side) (Piece, error)
}
