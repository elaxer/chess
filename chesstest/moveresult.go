package chesstest

import (
	"fmt"

	"github.com/elaxer/chess"
)

type MoveResultMock struct {
	MoveValue          chess.Move
	TurnValue          chess.Color
	CapturedPieceValue chess.Piece
	BoardNewStateValue chess.State
	StringValue        string
	ValidateFunc       func() error
}

func (m *MoveResultMock) Move() chess.Move {
	return m.MoveValue
}

func (m *MoveResultMock) Side() chess.Color {
	return m.TurnValue
}

func (m *MoveResultMock) CapturedPiece() chess.Piece {
	return m.CapturedPieceValue
}

func (m *MoveResultMock) SetBoardNewState(state chess.State) {
	m.BoardNewStateValue = state
}

func (m *MoveResultMock) BoardNewState() chess.State {
	return m.BoardNewStateValue
}

func (m *MoveResultMock) String() string {
	if m.StringValue != "" {
		return m.StringValue
	}

	return fmt.Sprintf(
		"Move: %v, Side: %v, Captured piece: %v, State: %v",
		m.MoveValue,
		m.TurnValue,
		m.CapturedPieceValue,
		m.BoardNewStateValue,
	)
}

func (m *MoveResultMock) Validate() error {
	if m.ValidateFunc != nil {
		return m.ValidateFunc()
	}

	return nil
}
