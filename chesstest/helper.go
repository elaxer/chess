package chesstest

import (
	"strings"
	"unicode"

	"github.com/elaxer/chess"
)

// NewPiece creates a new PieceMock instance based on the provided string representation.
// The string should be a single character representing the piece notation (e.g., "P", "N", "B", "r", "q", "k").
// The case of the character determines the color: uppercase for white, lowercase for black.
func NewPiece(str string) chess.Piece {
	if len(str) != 1 {
		panic("piece string must be a single character")
	}

	color := chess.ColorWhite
	if unicode.IsLower([]rune(str)[0]) {
		color = chess.ColorBlack
	}

	return &PieceMock{NotationValue: strings.ToUpper(str), ColorValue: color, StringValue: str}
}

// MoveStrings converts a list of move strings into a slice of chess.Move instances.
func MoveStrings(moveStrings ...string) []chess.Move {
	moves := make([]chess.Move, 0, len(moveStrings))
	for _, move := range moveStrings {
		moves = append(moves, chess.StringMove(move))
	}

	return moves
}
