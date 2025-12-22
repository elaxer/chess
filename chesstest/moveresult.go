package chesstest

import (
	"fmt"

	"github.com/elaxer/chess"
)

type MoveResultMock struct {
	MoveValue          chess.Move
	SideValue          chess.Side
	CapturedPieceValue chess.Piece
	BoardNewStateValue chess.State
	StringValue        string
	ValidateFunc       func() error
}

func (m *MoveResultMock) Move() chess.Move {
	return m.MoveValue
}

func (m *MoveResultMock) Side() chess.Side {
	return m.SideValue
}

func (m *MoveResultMock) CapturedPiece() chess.Piece {
	return m.CapturedPieceValue
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
		m.SideValue,
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
