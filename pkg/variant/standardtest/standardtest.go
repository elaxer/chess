package standardtest

import (
	"strings"
	"unicode"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/board"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
)

type Placement struct {
	Piece    chess.Piece
	Position position.Position
}

func NewEmpty(turn chess.Side, placements []Placement) chess.Board {
	board := board.NewFactory().CreateEmpty(turn)
	for _, placement := range placements {
		board.Squares().PlacePiece(placement.Piece, placement.Position)
	}

	return board
}

func MustNewFromMoves(moveStrings []string) chess.Board {
	board, err := NewFromMoves(moveStrings)
	if err != nil {
		panic(err)
	}

	return board
}

func NewFromMoves(moveStrings []string) (chess.Board, error) {
	moves := make([]chess.Move, 0, len(moveStrings))
	for _, notation := range moveStrings {
		moves = append(moves, chess.StringMove(notation))
	}

	return board.NewFactory().CreateFromMoves(moves)
}

// NewPiece creates a new piece by string.
// Created piece marked as not moved.
// P, R, N, B, Q, K - creates white piece
// p, r, n, b, q, k - creates black piece
func NewPiece(str string) chess.Piece {
	return newPiece(str, false)
}

// NewPieceW creates a new walked piece by string.
// See NewPiece for more details
func NewPieceW(str string) chess.Piece {
	return newPiece(str, true)
}

func newPiece(str string, isMoved bool) chess.Piece {
	side := chess.SideWhite
	if unicode.IsLower([]rune(str)[0]) {
		side = chess.SideBlack
	}

	notationUpper := strings.ToUpper(str)
	if notationUpper == "P" {
		notationUpper = ""
	}

	newPiece := piece.New(notationUpper, side)
	if isMoved {
		newPiece.MarkMoved()
	}

	return newPiece
}
