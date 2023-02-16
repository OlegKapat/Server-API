package apiserver

import "Go-http-rest-api/internal/api/store"

type Config struct {
	BindAdrr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAdrr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
