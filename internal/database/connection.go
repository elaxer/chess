package database

import (
	"context"
	"database/sql"
)

// Connection представляет собой интерфейс для работы с базой данных.
// Является адаптером для sql.DB и sql.Tx.
type Connection interface {
	Exec(query string, args ...any) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	Select(dest any, query string, args ...any) error
	SelectContext(ctx context.Context, dest any, query string, args ...any) error
	Get(dest any, query string, args ...any) error
	GetContext(ctx context.Context, dest any, query string, args ...any) error
}
