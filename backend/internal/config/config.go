package config

import "os"

type Config struct {
	Addr        string
	DatabaseURL string
}

func Load() Config {
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://fixit:secret@localhost:5438/fixit?sslmode=disable"
	}
	return Config{
		Addr:        addr,
		DatabaseURL: dsn,
	}
}
