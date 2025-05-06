package mover

import (
	"github.com/elaxer/chess/pkg/chess"
	mv "github.com/elaxer/chess/pkg/variant/standard/move"
)

var (
	normalMover    = new(Normal)
	castlingMover  = new(Castling)
	promotionMover = new(Promotion)
)

func MakeMove(move chess.Move, board chess.Board) (chess.Move, error) {
	switch m := move.(type) {
	case chess.RawMove:
		return nil, nil
	case *mv.Normal:
		return normalMover.Make(m, board)
	case *mv.Promotion:
		return promotionMover.Make(m, board)
	case *mv.Castling:
		return castlingMover.Make(m.CastlingType, board)

	}

	return nil, nil
}

func Undo(move chess.Move, board chess.Board) error {
	switch m := move.(type) {
	case *mv.Normal:
		return normalMover.Undo(m, board)
	}

	return nil
}

func modifyCheckMate(checkMate *mv.CheckMate, board chess.Board) {
	board.NextTurn()
	defer board.NextTurn()

	state := board.State()
	checkMate.IsCheck = state.IsCheck()
	checkMate.IsMate = state.IsMate()
}
