// Package chesstest provides utility functions and mocks for testing chess-related components.
package chesstest

import (
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/elaxer/chess"
)

type PieceMock struct {
	SideValue        chess.Side
	IsMovedValue     bool
	NotationValue    string
	WeightValue      uint8
	StringValue      string
	PseudoMovesValue chess.PositionSet
}

func (m *PieceMock) Side() chess.Side {
	return m.SideValue
}

func (m *PieceMock) IsMoved() bool {
	return m.IsMovedValue
}

func (m *PieceMock) SetIsMoved(isMoved bool) {
	m.IsMovedValue = isMoved
}

func (m *PieceMock) PseudoMoves(from chess.Position, squares *chess.Squares) chess.PositionSet {
	if m.PseudoMovesValue == nil {
		m.PseudoMovesValue = mapset.NewSet[chess.Position]()
	}

	return m.PseudoMovesValue
}

func (m *PieceMock) Notation() string {
	return m.NotationValue
}

func (m *PieceMock) Weight() uint8 {
	return m.WeightValue
}

func (m *PieceMock) String() string {
	if m.StringValue != "" || m.NotationValue == "" {
		return m.StringValue
	}

	notation := m.NotationValue

	if m.SideValue == chess.SideBlack {
		notation = strings.ToLower(notation)
	}

	return notation
}
