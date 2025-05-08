package standard

import (
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
)

type factory struct {
}

func NewFactory() chess.BoardFactory {
	return &factory{}
}

func (f *factory) CreateEmpty() chess.Board {
	squares := make(chess.Squares, 0, 64)
	for i := position.FileA; i <= position.FileH; i++ {
		for j := position.Rank1; j <= position.Rank8; j++ {
			squares = append(squares, &chess.Square{Position: position.New(i, j)})
		}
	}

	return &standard{
		turn:           chess.SideWhite,
		squares:        squares,
		movesHistory:   make([]chess.Move, 0, 128),
		capturedPieces: make([]chess.Piece, 0, 30),
	}
}

func (f *factory) CreateFilled() chess.Board {
	board := f.CreateEmpty()
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

	for i := position.FileA; i <= position.FileH; i++ {
		board.Squares().AddPiece(piece.New(notations[i-1], chess.SideWhite), position.New(i, position.Rank1))
		board.Squares().AddPiece(piece.NewPawn(chess.SideWhite), position.New(i, position.Rank2))

		board.Squares().AddPiece(piece.New(notations[i-1], chess.SideBlack), position.New(i, position.Rank8))
		board.Squares().AddPiece(piece.NewPawn(chess.SideBlack), position.New(i, position.Rank7))
	}

	return board
}

func (f *factory) CreateFromMoves(moves []chess.Move) (chess.Board, error) {
	board := f.CreateFilled()

	for _, move := range moves {
		if err := board.MakeMove(move); err != nil {
			return nil, err
		}
	}

	return board, nil
}
