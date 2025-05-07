package visualizer

import (
	"fmt"
	"io"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
)

func Visualize(squares chess.Squares, writer io.Writer) {
	for i := 1; i <= 8; i++ {
		for j := 1; j <= 8; j++ {
			square := squares.GetByPosition(position.New(position.File(j), position.Rank(i)))
			if square.IsEmpty() {
				fmt.Fprint(writer, ". ")

				continue
			}

			piece := square.Piece.Notation()
			if piece == chess.NotationPawn {
				piece = "P"
			}

			fmt.Fprintf(writer, "%s ", piece)
		}

		fmt.Fprintln(writer)
	}
}
