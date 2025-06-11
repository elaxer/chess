package chesstest

import (
	"github.com/elaxer/chess"
	"github.com/elaxer/chess/encoding/fen"
	"github.com/elaxer/chess/position"
)

// NewPiece creates a new PieceMock instance based on the provided string representation.
// The string should be a single character representing the piece notation (e.g., "P", "N", "B", "r", "q", "k").
// The case of the character determines the side: uppercase for white, lowercase for black.
func NewPiece(str string) chess.Piece {
	piece, err := new(PieceFactoryMock).CreateFromString(str)
	if err != nil {
		panic(err)
	}

	return piece
}

// NewPieceM creates a new PieceMock instance based on the provided string representation and marked as moved.
// The string should be a single character representing the piece notation (e.g., "P", "N", "B", "r", "q", "k").
// The case of the character determines the side: uppercase for white, lowercase for black.
func NewPieceM(str string) chess.Piece {
	piece := NewPiece(str)
	piece.MarkMoved()

	return piece
}

// BoardFromMoves creates a new chess.Board instance from the provided move strings.
// Provided move strings will be assigned to the board's moves history.
func BoardFromMoves(moveStrings ...string) chess.Board {
	board, err := new(FactoryMock).CreateFromMoves(MoveStrings(moveStrings...))
	if err != nil {
		panic(err)
	}

	return board
}

// MoveStrings converts a list of move strings into a slice of chess.Move instances.
func MoveStrings(moveStrings ...string) []chess.Move {
	moves := make([]chess.Move, 0, len(moveStrings))
	for _, move := range moveStrings {
		moves = append(moves, chess.StringMove(move))
	}

	return moves
}

// DecodeFEN8x8 decodes a FEN string into a chess.Board instance with an 8x8 edge position.
func DecodeFEN8x8(str string) chess.Board {
	return DecodeFEN(str, defaultEdgePosition)
}

// DecodeFEN decodes a FEN string into a chess.Board instance with the specified edge position.
func DecodeFEN(str string, edgePosition position.Position) chess.Board {
	decoder := fen.NewDecoder(&FactoryMock{EdgePosition: edgePosition}, new(PieceFactoryMock))
	board, err := decoder.Decode(str)
	if err != nil {
		panic(err)
	}

	return board
}
