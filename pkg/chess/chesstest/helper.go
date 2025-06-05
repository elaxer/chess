package chesstest

import (
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/encoding/fen"
	"github.com/elaxer/chess/pkg/chess/position"
)

// MustPieceFromString creates a new PieceMock instance based on the provided string representation.
// The string should be a single character representing the piece notation (e.g., "P", "N", "B", "r", "q", "k").
// The case of the character determines the side: uppercase for white, lowercase for black.
func MustPieceFromString(str string) chess.Piece {
	piece, err := new(PieceFactoryMock).CreateFromString(str)
	if err != nil {
		panic(err)
	}

	return piece
}

// MustPieceMFromString creates a new PieceMock instance based on the provided string representation and marked as moved.
// The string should be a single character representing the piece notation (e.g., "P", "N", "B", "r", "q", "k").
// The case of the character determines the side: uppercase for white, lowercase for black.
func MustPieceMFromString(str string) chess.Piece {
	piece := MustPieceFromString(str)
	piece.MarkMoved()

	return piece
}

func MustBoardFromMoveStrings(moveStrings []string) chess.Board {
	moves := make([]chess.Move, 0, len(moveStrings))
	for _, move := range moveStrings {
		moves = append(moves, chess.StringMove(move))
	}

	board, err := new(FactoryMock).CreateFromMoves(moves)
	if err != nil {
		panic(err)
	}

	return board
}

func MustDecodeFEN8x8(str string) chess.Board {
	return MustDecodeFEN(str, defaultEdgePosition)
}

func MustDecodeFEN(str string, edgePosition position.Position) chess.Board {
	decoder := &fen.Decoder{BoardFactory: &FactoryMock{EdgePosition: edgePosition}, PieceFactory: new(PieceFactoryMock)}
	board, err := decoder.Decode(str)
	if err != nil {
		panic(err)
	}

	return board
}
