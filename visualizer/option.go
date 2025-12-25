package visualizer

import "github.com/elaxer/chess/metric"

const (
	// OptionOrientationDefault displays board from white perspective.
	OptionOrientationDefault OptionOrientation = iota
	// OptionOrientationReversed displays board from black perspective.
	OptionOrientationReversed
	// OptionOrientationByTurn displays board from the side whose turn it is.
	OptionOrientationByTurn
)

// OptionOrientation represents how the board should be oriented when visualizing.
type OptionOrientation uint8

// Options holds settings for the Visualizer.
type Options struct {
	// Orientation defines how board ranks/files are shown.
	Orientation OptionOrientation
	// DisplayPositions enables printing rank and file indices around the board.
	DisplayPositions bool
	// MetricFuncs are optional metric functions to render below the board.
	MetricFuncs []metric.MetricFunc
}
