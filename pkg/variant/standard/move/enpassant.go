package move

import (
	"fmt"

	"github.com/elaxer/chess/pkg/chess/position"
)

type EnPassant struct {
	*CheckMate
	From, To position.Position
}

func (m *EnPassant) Notation() string {
	return fmt.Sprintf("%sx%s%s", m.From, m.To, m.CheckMate)
}

func (m *EnPassant) String() string {
	return m.Notation()
}
