package chesstest

import (
	"errors"
	"strings"
	"unicode"

	"github.com/elaxer/chess/pkg/chess"
)

type PieceFactoryMock struct {
	stringValue            string
	CreateFromNotationFunc func(notation string, side chess.Side) (chess.Piece, error)
	CreateFromStringFunc   func(str string) (chess.Piece, error)
}

func (f *PieceFactoryMock) CreateFromNotation(notation string, side chess.Side) (chess.Piece, error) {
	if f.CreateFromNotationFunc != nil {
		return f.CreateFromNotationFunc(notation, side)
	}

	return &PieceMock{NotationValue: notation, SideValue: side, StringValue: f.stringValue}, nil
}

func (f *PieceFactoryMock) CreateFromString(str string) (chess.Piece, error) {
	if f.CreateFromStringFunc != nil {
		return f.CreateFromStringFunc(str)
	}

	if len(str) != 1 {
		// todo
		return nil, errors.New("piece string must be a single character")
	}

	side := chess.SideWhite
	if unicode.IsLower([]rune(str)[0]) {
		side = chess.SideBlack
	}

	f.stringValue = str

	return f.CreateFromNotation(strings.ToUpper(str), side)
}
