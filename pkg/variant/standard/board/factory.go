package board

import (
	"fmt"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
	"github.com/elaxer/chess/pkg/variant/standard/state/rule"
)

var edgePosition = position.New(position.FileH, position.Rank8)

var firstRowPieceNotations = []string{
	piece.NotationRook,
	piece.NotationKnight,
	piece.NotationBishop,
	piece.NotationQueen,
	piece.NotationKing,
	piece.NotationBishop,
	piece.NotationKnight,
	piece.NotationRook,
}

var stateRules = []rule.Rule{
	rule.Mate,
	rule.Stalemate,
	rule.Check,

	rule.FiftyMove,
}

type factory struct {
}

func NewFactory() chess.BoardFactory {
	return &factory{}
}

func (f *factory) CreateEmpty(turn chess.Side) chess.Board {
	return &board{
		turn:           turn,
		squares:        chess.NewSquares(edgePosition),
		movesHistory:   make([]chess.Move, 0, 128),
		capturedPieces: make([]chess.Piece, 0, 30),

		stateRules: stateRules,
	}
}

func (f *factory) CreateFilled() chess.Board {
	board := f.CreateEmpty(chess.SideWhite)
	for i, notation := range firstRowPieceNotations {
		file := position.File(i + 1)

		board.Squares().PlacePiece(piece.New(notation, chess.SideWhite), position.New(file, position.Rank1))
		board.Squares().PlacePiece(piece.NewPawn(chess.SideWhite), position.New(file, position.Rank1+1))

		board.Squares().PlacePiece(piece.New(notation, chess.SideBlack), position.New(file, edgePosition.Rank))
		board.Squares().PlacePiece(piece.NewPawn(chess.SideBlack), position.New(file, edgePosition.Rank-1))
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
