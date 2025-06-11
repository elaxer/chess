package chesstest

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/elaxer/chess"
	"github.com/elaxer/chess/position"
)

type BoardMock struct {
	SquaresValue      *chess.Squares
	TurnValue         chess.Side
	StateFunc         func(side chess.Side) chess.State
	MovesHistoryValue []chess.MoveResult
	LegalMovesMap     map[chess.Piece]position.Set
	MakeMoveFunc      func(move chess.Move) (chess.MoveResult, error)
	UndoLastMoveFunc  func() (chess.MoveResult, error)
}

func (s *BoardMock) Squares() *chess.Squares {
	return s.SquaresValue
}

func (s *BoardMock) Turn() chess.Side {
	return s.TurnValue
}

func (s *BoardMock) State(side chess.Side) chess.State {
	if s.StateFunc != nil {
		return s.StateFunc(side)
	}

	return chess.StateClear
}

func (s *BoardMock) MovesHistory() []chess.MoveResult {
	return s.MovesHistoryValue
}

func (s *BoardMock) Moves(side chess.Side) position.Set {
	moves := mapset.NewSet[position.Position]()
	for _, piece := range s.Squares().GetAllPieces(side) {
		moves.Union(s.LegalMoves(piece))
	}

	return moves
}

func (s *BoardMock) LegalMoves(piece chess.Piece) position.Set {
	if piece == nil {
		return mapset.NewSet[position.Position]()
	}

	return piece.PseudoMoves(s.Squares().GetByPiece(piece), s.Squares())
}

func (s *BoardMock) MakeMove(move chess.Move) (chess.MoveResult, error) {
	if s.MakeMoveFunc != nil {
		return s.MakeMoveFunc(move)
	}

	result := &MoveResultMock{
		MoveValue:          move,
		SideValue:          s.TurnValue,
		BoardNewStateValue: s.State(!s.TurnValue),
		StringValue:        move.String(),
	}
	s.MovesHistoryValue = append(s.MovesHistoryValue, result)
	s.TurnValue = !s.TurnValue

	return result, nil
}

func (s *BoardMock) UndoLastMove() (chess.MoveResult, error) {
	if s.UndoLastMoveFunc != nil {
		return s.UndoLastMoveFunc()
	}

	if len(s.MovesHistoryValue) == 0 {
		return nil, nil
	}

	return s.MovesHistoryValue[len(s.MovesHistoryValue)-1], nil
}
