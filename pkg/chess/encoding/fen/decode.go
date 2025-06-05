package fen

import (
	"github.com/elaxer/chess/pkg/chess"
)

const regexp = "^([1-8PNBRQK]+/){7}[1-8PNBRQK]+"

func Decode(fen string, boardFactory chess.BoardFactory) (chess.Board, error) {
	return nil, nil
}
