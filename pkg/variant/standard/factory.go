package standard

import (
	"fmt"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
)

var edgePosition = position.New(position.FileH, position.Rank8)

type factory struct {
}

func NewFactory() chess.BoardFactory {
	return &factory{}
}

func (f *factory) CreateEmpty(turn chess.Side) chess.Board {
	return &standard{
		turn:           turn,
		squares:        chess.NewSquares(edgePosition),
		movesHistory:   make([]chess.Move, 0, 128),
		capturedPieces: make([]chess.Piece, 0, 30),
	}
}

func (f *factory) CreateFilled() chess.Board {
	board := f.CreateEmpty(chess.SideWhite)
	notations := []chess.PieceNotation{
		chess.NotationRook,
		chess.NotationKnight,
		chess.NotationBishop,
		chess.NotationQueen,
		chess.NotationKing,
		chess.NotationBishop,
		chess.NotationKnight,
		chess.NotationRook,
	}

	for i, notation := range notations {
		file := position.File(i + 1)

		board.Squares().AddPiece(piece.New(notation, chess.SideWhite), position.New(file, position.Rank1))
		board.Squares().AddPiece(piece.NewPawn(chess.SideWhite), position.New(file, position.Rank1+1))

		board.Squares().AddPiece(piece.New(notation, chess.SideBlack), position.New(file, edgePosition.Rank))
		board.Squares().AddPiece(piece.NewPawn(chess.SideBlack), position.New(file, edgePosition.Rank-1))
	}

	return board
}

func (f *factory) CreateFromMoves(moves []chess.Move) (chess.Board, error) {
	board := f.CreateFilled()
	for i, move := range moves {
		if err := board.MakeMove(move); err != nil {
			return nil, fmt.Errorf("%s#%d: %w", move, i+1, err)
		}
	}

	return board, nil
}
