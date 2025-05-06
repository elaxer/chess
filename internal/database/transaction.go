package database

import (
	"context"

	"github.com/elaxer/chess/internal/ctxkey"
	"github.com/jmoiron/sqlx"
)

// BeginTx начинает транзакцию и возвращает новый контекст с транзакцией.
func BeginTx(ctx context.Context, db *sqlx.DB) (*sqlx.Tx, context.Context, error) {
	tx, err := db.Beginx()

	if err != nil {
		return nil, nil, err
	}

	return tx, context.WithValue(ctx, ctxkey.Tx, tx), nil
}

// BeginTx начинает транзакцию и возвращает новый контекст с транзакцией.
// Если возникла ошибка, то вызывается panic.
func MustBeginTx(ctx context.Context, db *sqlx.DB) (*sqlx.Tx, context.Context) {
	tx, ctx, err := BeginTx(ctx, db)
	if err != nil {
		panic(err)
	}

	return tx, ctx
}

// TxOrDB возвращает транзакцию, если она есть в контексте, иначе возвращает базу данных.
func TxOrDB(ctx context.Context, db Connection) Connection {
	if tx, ok := ctx.Value(ctxkey.Tx).(*sqlx.Tx); ok {
		return tx
	}

	return db
}
