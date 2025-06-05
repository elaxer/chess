// Package chesstest provides utility functions and mocks for testing chess-related components.
package chesstest

import (
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
)

type PieceMock struct {
	SideValue        chess.Side
	IsMovedValue     bool
	NotationValue    string
	WeightValue      uint8
	StringValue      string
	PseudoMovesValue position.Set
}

func (m *PieceMock) Side() chess.Side {
	return m.SideValue
}

func (m *PieceMock) IsMoved() bool {
	return m.IsMovedValue
}

func (m *PieceMock) MarkMoved() {
	m.IsMovedValue = true
}

func (m *PieceMock) PseudoMoves(from position.Position, squares *chess.Squares) position.Set {
	if m.PseudoMovesValue == nil {
		m.PseudoMovesValue = mapset.NewSet[position.Position]()
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
