package main

import (
	"log"

	"github.com/MaxKut3/BudgetBot/config"
	"github.com/MaxKut3/BudgetBot/internal/app"
)

func main() {
	cfg, err := config.NewBotConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
