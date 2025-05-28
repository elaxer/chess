package standardtest

import (
	"strings"
	"unicode"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard"
	"github.com/elaxer/chess/pkg/variant/standard/move"
	resolver "github.com/elaxer/chess/pkg/variant/standard/moveresolver"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
)

type Placement struct {
	Piece    chess.Piece
	Position position.Position
}

func NewEmpty(turn chess.Side, placements []Placement) chess.Board {
	board := standard.NewFactory().CreateEmpty(turn)
	for _, placement := range placements {
		board.Squares().PlacePiece(placement.Piece, placement.Position)
	}

	return board
}

func NewFromMoves(notations []string) chess.Board {
	board, err := standard.NewFactory().CreateFromMoves(NotationsToMoves(notations))
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

func ResolveNormal(move *move.Normal, board chess.Board) *move.Normal {
	resolvedMove, err := resolver.ResolveNormal(move, board, board.Turn())
	if err != nil {
		panic(err)
	}

	return resolvedMove
}

func ResolvePromotion(move *move.Promotion, board chess.Board) *move.Promotion {
	move.Normal.PieceNotation = piece.NotationPawn
	move.Normal = ResolveNormal(move.Normal, board)

	return move
}

func NewPiece(notation string, isMoved bool) chess.Piece {
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
