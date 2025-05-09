package visualizer

import (
	"fmt"
	"io"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
)

func Visualize(squares chess.Squares, writer io.Writer) {
	for rank := range squares.EdgePosition().Rank {
		for file := range squares.EdgePosition().File {
			square := squares.GetByPosition(position.New(file+1, rank+1))
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
