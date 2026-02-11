package chesstest

import (
	"fmt"

	"github.com/elaxer/chess"
)

type MoveMock struct {
	InputValue         string
	TurnValue          chess.Color
	CapturedPieceValue chess.Piece
	BoardNewStateValue chess.State
	StringValue        string
	ValidateFunc       func() error
}

func (m *MoveMock) Input() string {
	return m.InputValue
}

func (m *MoveMock) Side() chess.Color {
	return m.TurnValue
}

func (m *MoveMock) CapturedPiece() chess.Piece {
	return m.CapturedPieceValue
}

func (m *MoveMock) SetBoardNewState(state chess.State) {
	m.BoardNewStateValue = state
}

func (m *MoveMock) BoardNewState() chess.State {
	return m.BoardNewStateValue
}

func (m *MoveMock) String() string {
	if m.StringValue != "" {
		return m.StringValue
	}

	return fmt.Sprintf(
		"Move: %v, Side: %v, Captured piece: %v, State: %v",
		m.InputValue,
		m.TurnValue,
		m.CapturedPieceValue,
		m.BoardNewStateValue,
	)
}

func (m *MoveMock) Validate() error {
	if m.ValidateFunc != nil {
		return m.ValidateFunc()
	}

	return nil
}
