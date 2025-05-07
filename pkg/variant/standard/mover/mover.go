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
	return nil, nil
}

func MakeMoveFromNotation(notation string, board chess.Board) (chess.Move, error) {
	return nil, nil
}

func modifyCheckMate(checkMate *mv.CheckMate, board chess.Board) {
	board.NextTurn()
	defer board.NextTurn()

	state := board.State()
	checkMate.IsCheck = state.IsCheck()
	checkMate.IsMate = state.IsMate()
}
