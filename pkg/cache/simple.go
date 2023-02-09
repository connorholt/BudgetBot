package cache

import (
	"sync"
	"time"
)

type CurrencyCache interface {
	Get(key string) (int, error)
	Set(key string, val int)
}

type Currency string
type Record struct {
	Value     int
	CreatedAt time.Time
}

type SimpleCache struct {
	mu sync.RWMutex
	mp map[Currency]Record
}

func (s *SimpleCache) Get(key string) (int, error) {
	return 0, nil
}

func (s *SimpleCache) Set(key string, val int) {
	return
}

func (s *SimpleCache) watch() {
	// for map
}

func NewSimple() CurrencyCache {
	cache := &SimpleCache{
		mu: sync.RWMutex{},
		mp: make(map[Currency]Record),
	}

	go cache.watch()

	return cache
}
