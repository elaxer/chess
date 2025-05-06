package config

type Websocket struct {
	Port    int `env:"PORT"`
	Origins []string
}
