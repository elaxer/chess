package standardtest

import (
	"strings"
	"unicode"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/board"
	"github.com/elaxer/chess/pkg/variant/standard/move/move"
	"github.com/elaxer/chess/pkg/variant/standard/move/resolver"
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

func NewFromMoves(notations []string) chess.Board {
	board, err := board.NewFactory().CreateFromMoves(NotationsToMoves(notations))
	if err != nil {
		panic(err)
	}

	return board
}

func NotationsToMoves(notations []string) []chess.Move {
	moves := make([]chess.Move, 0, len(notations))
	for _, notation := range notations {
		moves = append(moves, chess.RawMove(notation))
	}

	return moves
}

func ResolveNormal(move *move.Normal, board chess.Board) {
	resolvedFrom, err := resolver.ResolveFrom(move, board, board.Turn())
	if err != nil {
		panic(err)
	}

	move.From = resolvedFrom
}

func ResolvePromotion(move *move.Promotion, board chess.Board) {
	move.Normal.PieceNotation = piece.NotationPawn
	ResolveNormal(move.Normal, board)
}

// NewPiece creates a new piece by notation.
// Created piece marked as not moved.
// P, R, N, B, Q, K - creates white piece
// p, r, n, b, q, k - creates black piece
func NewPiece(notation string) chess.Piece {
	return newPiece(notation, false)
}

// NewPieceW creates a new walked piece by notation.
// See NewPiece for more details
func NewPieceW(notation string) chess.Piece {
	return newPiece(notation, true)
}

func newPiece(notation string, isMoved bool) chess.Piece {
	side := chess.SideWhite
	if unicode.IsLower([]rune(notation)[0]) {
		side = chess.SideBlack
	}

	notationUpper := strings.ToUpper(notation)
	if notationUpper == "P" {
		notationUpper = ""
	}

	newPiece := piece.New(notationUpper, side)
	if isMoved {
		newPiece.MarkMoved()
	}

	return newPiece
}
