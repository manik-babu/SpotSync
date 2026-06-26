package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Port      string
	Dsn       string
	JwtSecret string
}

func LoadEnv() *Env {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Failed to load environment variables")
	}
	return &Env{
		Port:      os.Getenv("PORT"),
		Dsn:       os.Getenv("DSN"),
		JwtSecret: os.Getenv("JWT_SECRET"),
	}
}
