package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Port string
}

func LoadEnv() *Env {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Failed to load environment variables")
	}
	return &Env{
		Port: os.Getenv("PORT"),
	}
}
