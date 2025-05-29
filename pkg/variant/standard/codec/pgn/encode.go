package pgn

import (
	"fmt"

	"github.com/elaxer/chess/pkg/chess"
)

func Encode(board chess.Board, tags map[string]string) string {
	pgn := ""
	for tag, value := range tags {
		pgn += fmt.Sprintf("[%s \"%s\"]\n", tag, value)
	}

	pgn += "\n"

	currentMoveNumber := 0
	for i, move := range board.MovesHistory() {
		if moveNumber := (i + 2) / 2; moveNumber != currentMoveNumber {
			currentMoveNumber = moveNumber
			pgn += fmt.Sprintf("%d. ", currentMoveNumber)
		}

		pgn += fmt.Sprintf("%s ", move)
	}

	return pgn + pgnResult(board)
}

func pgnResult(board chess.Board) string {
	if !board.State(board.Turn()).Type().IsTerminal() {
		return "*"
	}
	if board.State(board.Turn()).Type().IsDraw() {
		return "1/2-1/2"
	}

	if board.Turn().IsBlack() {
		return "0-1"
	}

	return "1-0"
}
