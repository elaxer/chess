package standard

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	mv "github.com/elaxer/chess/pkg/variant/standard/move"
	"github.com/elaxer/chess/pkg/variant/standard/move/validator"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
)

// EncodeFEN encodes the board position in FEN format.
// The FEN format is a standard notation for describing chess positions.
// It consists of six fields separated by spaces:
// 1. Piece placement
// 2. Active color
// 3. Castling availability
// 4. En passant target square
// 5. Halfmove clock
// 6. Fullmove number
func EncodeFEN(board chess.Board) string {
	return fmt.Sprintf(
		"%s %s %s %s %s %d",
		fenPosition(board.Squares()),
		board.Turn(),
		fenCastlings(board),
		fenEnPassantPosition(board),
		fenHalfmoveClock(board),
		len(board.MovesHistory())/2+1,
	)
}

func fenPosition(squares chess.Squares) string {
	fen := ""
	for i := range squares.EdgePosition().Rank {
		row := ""
		emptySquares := 0
		for j := range squares.EdgePosition().File {
			p, _ := squares.GetByPosition(position.New(j+1, i+1))
			if p == nil {
				emptySquares++

				continue
			}

			if emptySquares > 0 {
				row += strconv.Itoa(emptySquares)
			}
			emptySquares = 0

			row += fmt.Sprintf("%s", p)
		}

		if emptySquares > 0 {
			row += strconv.Itoa(emptySquares)
		}

		fen += row + "/"
	}

	return fen[:len(fen)-1]
}

func fenCastlings(board chess.Board) string {
	castlings := map[string]bool{
		"K": validator.ValidateCastling(mv.CastlingShort, chess.SideWhite, board) == nil,
		"Q": validator.ValidateCastling(mv.CastlingLong, chess.SideWhite, board) == nil,
		"k": validator.ValidateCastling(mv.CastlingShort, chess.SideBlack, board) == nil,
		"q": validator.ValidateCastling(mv.CastlingLong, chess.SideBlack, board) == nil,
	}

	fen := ""
	for castling, isValid := range castlings {
		if isValid {
			fen += castling
		}
	}

	if fen == "" {
		return "-"
	}

	return fen
}

func fenEnPassantPosition(board chess.Board) string {
	if len(board.MovesHistory()) == 0 {
		return "-"
	}

	lastMove := board.MovesHistory()[len(board.MovesHistory())-1]
	normalMove, ok := lastMove.(*mv.Normal)
	if !ok || normalMove.PieceNotation != piece.NotationPawn {
		return "-"
	}

	if normalMove.To.Rank != normalMove.From.Rank+(piece.PawnRankDirection(!board.Turn())*2) {
		return "-"
	}

	passant := position.New(
		normalMove.To.File,
		normalMove.From.Rank+piece.PawnRankDirection(!board.Turn()),
	)

	return passant.String()
}

func fenHalfmoveClock(board chess.Board) string {
	moves := slices.Clone(board.MovesHistory())
	slices.Reverse(moves)

	count := 0
	for _, move := range moves {
		normalMove, ok := move.(*mv.Normal)
		if !ok || normalMove.PieceNotation == piece.NotationPawn || normalMove.IsCapture {
			count = 0

			continue
		}

		count++
	}

	return strconv.Itoa(count/2 + 1)
}
