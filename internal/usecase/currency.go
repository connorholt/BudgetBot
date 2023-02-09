package usecase

import (
	"context"
	"github.com/MaxKut3/BudgetBot/config"
)

type CurrencyConverter interface {
	GetValue() int
}

type currencyConverter struct {
	cfg *config.Config
}

func NewCurrencyConverter(ctx context.Context, cfg *config.Config) CurrencyConverter {
	return &currencyConverter{
		cfg: cfg,
	}
}

func (c *currencyConverter) GetValue() int {
	providers := c.cfg.Providers

	for _, provideFn := range providers {
		go provideFn()
	}

	return 0
}
