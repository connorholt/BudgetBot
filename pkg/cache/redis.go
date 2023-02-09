package cache

import (
	"sync"
)

type RedisCache struct {
	mu sync.RWMutex
	mp map[Currency]Record
}

func (s *RedisCache) Get(key string) (int, error) {
	return 0, nil
}

func (s *RedisCache) Set(key string, val int) {
	return
}

func NewRedis() CurrencyCache {
	cache := &RedisCache{
		mu: sync.RWMutex{},
		mp: make(map[Currency]Record),
	}

	return cache
}
