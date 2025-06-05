package chesstest

type MoveMock struct {
	ValidateFunc func() error
	StringValue  string
}

func (m *MoveMock) Validate() error {
	if m.ValidateFunc != nil {
		return m.ValidateFunc()
	}

	return nil
}

func (m *MoveMock) String() string {
	return m.StringValue
}
