package config

import "fmt"

type DB struct {
	Host     string `env:"HOST"`
	Port     int    `env:"PORT"`
	SSLMode  string `env:"SSL_MODE"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	DBName   string `env:"DB_NAME"`
}

func (db *DB) String() string {
	return fmt.Sprintf(
		"host=%s port=%d sslmode=%s user=%s password=%s dbname=%s",
		db.Host,
		db.Port,
		db.SSLMode,
		db.User,
		db.Password,
		db.DBName,
	)
}
