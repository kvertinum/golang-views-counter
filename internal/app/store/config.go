package store

type RedisConfig struct {
	RedisURL string `toml:"redis_url"`
	Password string `toml:"password"`
}
