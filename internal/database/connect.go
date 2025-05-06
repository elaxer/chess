package database

import (
	"github.com/elaxer/chess/internal/config"
	"github.com/jmoiron/sqlx"
)

// ConnectAndPing подключается к базе данных и проверяет соединение.
// Если соединение не установлено, то возвращает ошибку.
func ConnectAndPing(dbConfig *config.DB) (*sqlx.DB, error) {
	db, err := sqlx.Open(driver, dbConfig.String())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// MustConnectAndPing подключается к базе данных и проверяет соединение.
// Если соединение не установлено, то вызывает panic.
func MustConnectAndPing(dbConfig *config.DB) *sqlx.DB {
	db, err := ConnectAndPing(dbConfig)

	if err != nil {
		panic(err)
	}

	return db
}
