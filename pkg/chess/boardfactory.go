package chess

type BoardFactory interface {
	CreateEmpty(turn Side) Board
	CreateFilled() Board
	CreateFromMoves(moves []Move) (Board, error)
}
