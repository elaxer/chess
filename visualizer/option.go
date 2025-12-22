package visualizer

import "github.com/elaxer/chess/metric"

const (
	OptionOrientationDefault OptionOrientation = iota
	OptionOrientationReversed
	OptionOrientationByTurn
)

type OptionOrientation uint8

type Options struct {
	Orientation      OptionOrientation
	DisplayPositions bool
	MetricFuncs      []metric.MetricFunc
}
