package visualizer

import "github.com/elaxer/chess/pkg/chess/metric"

type OptionOrientation uint8

const (
	OptionOrientationDefault OptionOrientation = iota
	OptionOrientationReversed
	OptionOrientationByTurn
)

type Options struct {
	Orientation OptionOrientation
	Positions   bool
	MetricFuncs []metric.MetricFunc
}
