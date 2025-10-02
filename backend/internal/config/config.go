package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Addr        string
	DatabaseURL string
	JWTSecret   string
	JWTTTL      time.Duration
}

func Load() Config {
	_ = godotenv.Load(".env")
	addr := os.Getenv("ADDR")
	dsn := os.Getenv("DATABASE_URL")
	secret := os.Getenv("JWT_SECRET")
	ttlenv := os.Getenv("JWT_TTL")
	ttl := 60 * time.Minute
	if ttlenv != "" {
		if m, ok := strconv.Atoi(ttlenv); ok == nil {
			ttl = time.Duration(m) * time.Minute
		}

	}

	return Config{
		Addr:        addr,
		DatabaseURL: dsn,
		JWTSecret:   secret,
		JWTTTL:      ttl,
	}

}
