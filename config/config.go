package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
}

type Config struct {
	BotApiKet string
	Providers map[string]func()
}

func NewBotConfig() (*Config, error) {
	botApiKey := os.Getenv("KEY")
	return &Config{
		BotApiKet: botApiKey,
		Providers: make(map[string]func()),
	}, nil
}
