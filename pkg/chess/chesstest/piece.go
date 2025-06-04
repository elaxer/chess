// Package chesstest provides utility functions and mocks for testing chess-related components.
package chesstest

import (
	"strings"
	"unicode"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
)

type PieceMock struct {
	SideValue        chess.Side
	IsMovedValue     bool
	NotationValue    string
	WeightValue      uint8
	PseudoMovesValue position.Set
}

// NewPieceMock creates a new PieceMock instance based on the provided string representation.
// The string should be a single character representing the piece notation (e.g., "P", "N", "B", "r", "q", "k").
// The case of the character determines the side: uppercase for white, lowercase for black.
func NewPieceMock(str string) chess.Piece {
	side := chess.SideWhite
	if unicode.IsLower([]rune(str)[0]) {
		side = chess.SideBlack
	}

	notation := strings.ToUpper(str)
	if notation == "P" {
		notation = ""
	}

	return &PieceMock{
		SideValue:        side,
		NotationValue:    notation,
		PseudoMovesValue: mapset.NewSet[position.Position](),
	}
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
	return m.NotationValue
}
