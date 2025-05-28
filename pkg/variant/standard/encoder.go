package standard

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/metric"
	mv "github.com/elaxer/chess/pkg/variant/standard/move"
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
		metric.CastlingsString(board).Value(),
		metric.EnPassantPosition(board).Value(),
		fenHalfmoveClock(board),
		len(board.MovesHistory())/2+1,
	)
}

func fenPosition(squares *chess.Squares) string {
	fen := ""
	for rank := range squares.EdgePosition().Rank {
		row := ""
		emptySquares := 0
		for file := range squares.EdgePosition().File {
			position, _ := squares.GetByPosition(position.New(file+1, rank+1))
			if position == nil {
				emptySquares++

				continue
			}

			if emptySquares > 0 {
				row += strconv.Itoa(emptySquares)
			}
			emptySquares = 0

			row += fmt.Sprintf("%s", position)
		}

		if emptySquares > 0 {
			row += strconv.Itoa(emptySquares)
		}

		fen += row + "/"
	}

	return fen[:len(fen)-1]
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
