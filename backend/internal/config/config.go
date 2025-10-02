package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Addr        string
	DatabaseURL string
}

func Load() Config {
	_ = godotenv.Load(".env")
	addr := os.Getenv("ADDR")
	dsn := os.Getenv("DATABASE_URL")
	return Config{
		Addr:        addr,
		DatabaseURL: dsn,
	}
}
