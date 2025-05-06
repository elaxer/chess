package parser

import (
	"fmt"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/variant/standard/move"
)

// ParseMove парсит нотацию и создает ход на основе нотации.
func ParseNormalMove(notation string, side chess.Side, board chess.Board) (*move.Normal, error) {
	move, err := move.NewNormal(notation)
	if err != nil {
		return nil, err
	}

	if err := move.From.Validate(); err == nil {
		return move, nil
	}

	pieces := make([]chess.Piece, 0, 8)
	for _, piece := range board.Squares().GetPieces(move.PieceNotation, side) {
		if piece.Moves(board).Has(move.To) {
			pieces = append(pieces, piece)
		}
	}

	if len(pieces) == 0 {
		return nil, fmt.Errorf("%w: не найдено подходящих фигур для хода", Err)
	}
	if len(pieces) == 1 {
		move.From = board.Squares().GetByPiece(pieces[0]).Position
	}

	for _, piece := range pieces {
		pos := board.Squares().GetByPiece(piece).Position
		if move.From.Rank == 0 && pos.File == move.From.File {
			move.From.Rank = pos.Rank
		}
		if move.From.File == 0 && pos.Rank == move.From.Rank {
			move.From.File = pos.File
		}
	}

	return move, nil
}
