package main

import (
	"os"

	"github.com/elaxer/chess/pkg/variant/standard"
	"github.com/elaxer/chess/pkg/visualizer"
)

func main() {
	squares := standard.NewFactory().CreateFilled().Squares()

	visualizer.Visualize(squares, os.Stdout)
}
