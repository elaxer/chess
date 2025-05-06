package main

import (
	"os"

	"github.com/elaxer/chess/pkg/variant/standard"
	"github.com/elaxer/chess/pkg/visualizer"
)

func main() {
	board := standard.NewBoardFactory().CreateFilled()

	visualizer.Visualize(board, os.Stdout)
}
