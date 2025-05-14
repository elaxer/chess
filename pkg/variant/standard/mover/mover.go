package mover

import (
	"errors"
	"fmt"

	"github.com/elaxer/chess/pkg/chess"
	mv "github.com/elaxer/chess/pkg/variant/standard/move"
)

var Err = errors.New("mover error")

var (
	normalMover    = new(Normal)
	castlingMover  = new(Castling)
	promotionMover = new(Promotion)
)

func MakeMove(move chess.Move, board chess.Board) (chess.Move, error) {
	return makeMoveFromNotation(move.Notation(), board)
}

// todo
func makeMoveFromNotation(notation string, board chess.Board) (chess.Move, error) {
	if move, err := mv.NewNormal(notation); err == nil {
		return normalMover.Make(move, board)
	}
	if move, err := mv.NewPromotion(notation); err == nil {
		return promotionMover.Make(move, board)
	}
	if move, err := mv.NewCastling(notation); err == nil {
		return castlingMover.Make(move.CastlingType, board)
	}

	return nil, fmt.Errorf("%w: invalid move", Err)
}
