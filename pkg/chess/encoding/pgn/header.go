package pgn

import "fmt"

type Header struct {
	Name  string
	Value string
}

func NewHeader(name, value string) Header {
	return Header{name, value}
}

func (h Header) String() string {
	return fmt.Sprintf("[%s \"%s\"]", h.Name, h.Value)
}
