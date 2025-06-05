package fen

import (
	"github.com/elaxer/chess/pkg/chess"
)

const regexp = "^(((1[0-6]|[1-9])|[PNBRQK])+/){15}((1[0-6]|[1-9])|[PNBRQK])+"

func Decode(fen string, boardFactory chess.BoardFactory) (chess.Board, error) {
	return nil, nil
}
