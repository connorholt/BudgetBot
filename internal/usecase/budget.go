package usecase

import (
	"context"
	"fmt"
)

type CurrencyCache interface {
	Get(key string) (int, error)
	Set(key string, val int)
}

type Handler interface {
	Create()
	List()
}

type handler struct {
	currencyConverter CurrencyConverter
	cache             CurrencyCache
}

func (h handler) Create() {
	// todo
	// value := CurrencyConverter(CacheConverter("key"))
	value, err := h.cache.Get("key")
	if err != nil {
		value := h.currencyConverter.GetValue()
		h.cache.Set("ket", value)
	}

	fmt.Println(value)
	return
}

func (h handler) List() {

}

func NewBudget(ctx context.Context, currencyConverter CurrencyConverter, cache CurrencyCache) Handler {
	return &handler{
		currencyConverter: currencyConverter,
		cache:             cache,
	}
}
