// Package metric provides utilities to compute various numeric and descriptive
// metrics for a chess board (for example: material balance, move counters,
// and last move). Metrics implement the Metric interface and can be used by
// encoders or visualizers to include additional information about a board.
package metric

import (
	"fmt"

	"github.com/elaxer/chess"
)

// MetricFunc is a function that produces a Metric for the given board.
// It can return nil if the metric is not applicable for the current position.
type MetricFunc func(board chess.Board) Metric

// Metric represents a named metric value produced from a board (for debugging or display).
// It provides a human-readable name and the underlying value.
type Metric interface {
	// Name returns the human-readable name of the metric.
	Name() string
	// Value returns the underlying value of the metric.
	Value() any
}

// metric is an internal implementation of the Metric interface.
type metric struct {
	name  string
	value any
}

// New creates a Metric with the given name and value.
func New(name string, value any) Metric {
	return &metric{name, value}
}

func (m *metric) Name() string {
	return m.name
}

func (m *metric) Value() any {
	return m.value
}

func (m *metric) String() string {
	return fmt.Sprintf("%s: %v", m.name, m.value)
}
