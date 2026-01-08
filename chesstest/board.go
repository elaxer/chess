package chesstest

import (
	"slices"

	"github.com/elaxer/chess"
)

type BoardMock struct {
	SquaresValue        *chess.Squares
	TurnValue           chess.Color
	StateFunc           func() chess.State
	MovesHistoryValue   []chess.MoveResult
	CapturedPiecesValue []chess.Piece
	LegalMovesMap       map[chess.Piece][]chess.Position
	MakeMoveFunc        func(move chess.Move) (chess.MoveResult, error)
	UndoLastMoveFunc    func() (chess.MoveResult, error)
}

func (s *BoardMock) Squares() *chess.Squares {
	return s.SquaresValue
}

func (s *BoardMock) Turn() chess.Color {
	return s.TurnValue
}

func (s *BoardMock) State() chess.State {
	if s.StateFunc != nil {
		return s.StateFunc()
	}

	return chess.StateClear
}

func (s *BoardMock) CapturedPieces() []chess.Piece {
	return s.CapturedPiecesValue
}

func (s *BoardMock) IsSquareAttacked(position chess.Position) bool {
	for piece := range s.Squares().GetAllPieces(!s.TurnValue) {
		from := s.SquaresValue.GetByPiece(piece)
		if slices.Contains(piece.PseudoMoves(from, s.SquaresValue), position) {
			return true
		}
	}

	return false
}

func (s *BoardMock) MoveHistory() []chess.MoveResult {
	return s.MovesHistoryValue
}

func (s *BoardMock) Moves() []chess.Position {
	moves := make([]chess.Position, 0, 128)
	for piece := range s.Squares().GetAllPieces(s.TurnValue) {
		moves = append(moves, s.LegalMoves(piece)...)
	}

	return moves
}

func (s *BoardMock) LegalMoves(piece chess.Piece) []chess.Position {
	if piece == nil {
		return make([]chess.Position, 0, 16)
	}

	return piece.PseudoMoves(s.Squares().GetByPiece(piece), s.Squares())
}

func (s *BoardMock) MakeMove(move chess.Move) (chess.MoveResult, error) {
	if s.MakeMoveFunc != nil {
		return s.MakeMoveFunc(move)
	}

	result := &MoveResultMock{
		MoveValue:          move,
		TurnValue:          s.TurnValue,
		BoardNewStateValue: s.State(),
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
