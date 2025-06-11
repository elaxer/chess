package chesstest

import (
	"fmt"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
)

var defaultEdgePosition = position.New(position.FileH, position.Rank8)

type FactoryMock struct {
	EdgePosition position.Position
	PieceFactory chess.PieceFactory

	CreateFunc          func(turn chess.Side, placement map[position.Position]chess.Piece) (chess.Board, error)
	CreateFilledFunc    func() chess.Board
	CreateFromMovesFunc func(moves []chess.Move) (chess.Board, error)
}

func (f *FactoryMock) Create(turn chess.Side, placement map[position.Position]chess.Piece) (chess.Board, error) {
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
		file := position.File(i + 1)

		if f.PieceFactory == nil {
			f.PieceFactory = &PieceFactoryMock{}
		}

		wP, _ := f.PieceFactory.CreateFromNotation(notation, chess.SideWhite)
		wPawn, _ := f.PieceFactory.CreateFromNotation(piece.NotationPawn, chess.SideWhite)
		board.Squares().PlacePiece(wP, position.New(file, position.RankMin))
		board.Squares().PlacePiece(wPawn, position.New(file, position.RankMin+1))

		bP, _ := f.PieceFactory.CreateFromNotation(notation, chess.SideBlack)
		bPawn, _ := f.PieceFactory.CreateFromNotation(piece.NotationPawn, chess.SideBlack)
		board.Squares().PlacePiece(bP, position.New(file, f.EdgePosition.Rank))
		board.Squares().PlacePiece(bPawn, position.New(file, f.EdgePosition.Rank-1))
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
