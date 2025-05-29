package move

type CastlingType bool

const (
	CastlingShort CastlingType = true
	CastlingLong  CastlingType = false
)

func (m CastlingType) IsShort() bool {
	return m == CastlingShort
}

func (m CastlingType) IsLong() bool {
	return m == CastlingLong
}

func (m CastlingType) String() string {
	return map[CastlingType]string{
		CastlingShort: "0-0",
		CastlingLong:  "0-0-0",
	}[m]
}
