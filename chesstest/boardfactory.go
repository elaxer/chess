package chesstest

import (
	"fmt"

	"github.com/elaxer/chess"
	"github.com/elaxer/chess/position"
)

var defaultEdgePosition = position.New(position.FileH, position.Rank8)

type FactoryMock struct {
	EdgePosition position.Position
	PieceFactory chess.PieceFactory

	CreateFunc          func(turn chess.Side, placement map[position.Position]chess.Piece) (chess.Board, error)
	CreateFilledFunc    func() chess.Board
	CreateFromMovesFunc func(moves []chess.Move) (chess.Board, error)
}

func (f *FactoryMock) Create(
	turn chess.Side,
	placement map[position.Position]chess.Piece,
) (chess.Board, error) {
	if f.CreateFunc != nil {
		return f.CreateFunc(turn, placement)
	}

	if f.EdgePosition.IsEmpty() {
		f.EdgePosition = defaultEdgePosition
	}

	squares, err := chess.SquaresFromPlacement(f.EdgePosition, placement)
	if err != nil {
		return nil, err
	}

	return &BoardMock{
		TurnValue:         turn,
		SquaresValue:      squares,
		MovesHistoryValue: make([]chess.MoveResult, 0, 128),
	}, nil
}

func (f *FactoryMock) CreateFilled() chess.Board {
	if f.CreateFilledFunc != nil {
		return f.CreateFilledFunc()
	}

	board, _ := f.Create(chess.SideWhite, nil)
	for i, notation := range []string{"R", "N", "B", "Q", "K", "B", "N", "R"} {
		//nolint:gosec
		file := position.File(i + 1)

		if f.PieceFactory == nil {
			f.PieceFactory = &PieceFactoryMock{}
		}

		wP, _ := f.PieceFactory.CreateFromNotation(notation, chess.SideWhite)
		wPawn, _ := f.PieceFactory.CreateFromString("P")
		must(board.Squares().PlacePiece(wP, position.New(file, position.RankMin)))
		must(board.Squares().PlacePiece(wPawn, position.New(file, position.RankMin+1)))

		bP, _ := f.PieceFactory.CreateFromNotation(notation, chess.SideBlack)
		bPawn, _ := f.PieceFactory.CreateFromString("p")
		must(board.Squares().PlacePiece(bP, position.New(file, f.EdgePosition.Rank)))
		must(board.Squares().PlacePiece(bPawn, position.New(file, f.EdgePosition.Rank-1)))
	}

	return board
}

func (f *FactoryMock) CreateFromMoves(moves []chess.Move) (chess.Board, error) {
	if f.CreateFromMovesFunc != nil {
		return f.CreateFromMovesFunc(moves)
	}

	board := f.CreateFilled()
	for i, move := range moves {
		if _, err := board.MakeMove(move); err != nil {
			return nil, fmt.Errorf("%s#%d: %w", move, i+1, err)
		}
	}

	return board, nil
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
