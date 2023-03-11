package server

import "github.com/Kvertinum01/views-counter/internal/app/store"

type Config struct {
	ServerAddr  string             `toml:"addr"`
	RedisConfig *store.RedisConfig `toml:"store"`
}
