package ctxkey

type ctxKey int

const (
	User ctxKey = iota

	Tx
)
