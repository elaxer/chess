package pgn

import (
	"fmt"
	"strings"

	"github.com/elaxer/chess/pkg/chess"
)

func Encode(board chess.Board, headers []Header) string {
	var pgn strings.Builder
	fmt.Fprintln(&pgn, EncodeHeaders(headers))

	movesStr := wrapText(EncodeMoves(board.MovesHistory()), 79)
	fmt.Fprint(&pgn, movesStr)

	return pgn.String() + " " + result(board)
}

func EncodeHeaders(headers []Header) string {
	var str strings.Builder
	for _, header := range headers {
		fmt.Fprintln(&str, header)
	}

	return str.String()
}

func EncodeMoves(moves []chess.MoveResult) string {
	var str strings.Builder
	currentMoveNumber := 0
	for i, move := range moves {
		if moveNumber := (i + 2) / 2; moveNumber != currentMoveNumber {
			currentMoveNumber = moveNumber
			fmt.Fprintf(&str, "%d. ", currentMoveNumber)
		}

		fmt.Fprintf(&str, "%s ", move)
	}

	return str.String()
}

func wrapText(text string, maxWidth int) string {
	if maxWidth <= 0 {
		return text
	}

	var result strings.Builder
	var lineLen int

	words := strings.Fields(text)

	for i, word := range words {
		if lineLen+len(word) > maxWidth {
			result.WriteString("\n")
			lineLen = 0
		} else if i != 0 {
			result.WriteString(" ")
			lineLen++
		}

		result.WriteString(word)
		lineLen += len(word)
	}

	return result.String()
}

func result(board chess.Board) string {
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
