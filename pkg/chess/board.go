package chess

import "github.com/elaxer/chess/pkg/chess/position"

type Board interface {
	Squares() *Squares
	Turn() Side
	State(side Side) State
	MovesHistory() []Move
	Moves(side Side) position.Set
	LegalMoves(piece Piece) position.Set
	MakeMove(move Move) error
}
