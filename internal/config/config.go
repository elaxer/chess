package config

import "github.com/caarlos0/env/v11"

type Config struct {
	DB         *DB        `env:",init" envPrefix:"DB_"`
	Websocket  *Websocket `env:",init" envPrefix:"WEBSOCKET_"`
	ServerPort int        `env:"SERVER_PORT"`
}

// Load загружает конфигурацию из переменных окружения.
func Load() (*Config, error) {
	cfg, err := env.ParseAsWithOptions[Config](env.Options{RequiredIfNoDef: true})
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

// MustLoad загружает конфигурацию из переменных окружения и паникует в случае ошибки.
func MustLoad() *Config {
	cfg, err := Load()
	if err != nil {
		panic(err)
	}

	return cfg
}
