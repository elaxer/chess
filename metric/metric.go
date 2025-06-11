package metric

import (
	"fmt"

	"github.com/elaxer/chess"
)

type MetricFunc func(board chess.Board) Metric

type Metric interface {
	Description() string
	Value() any
}

type metric struct {
	description string
	value       any
}

func New(description string, value any) Metric {
	return &metric{description, value}
}

func (m *metric) Description() string {
	return m.description
}

func (m *metric) Value() any {
	return m.value
}

func (m *metric) String() string {
	return fmt.Sprintf("%s: %v", m.description, m.value)
}
