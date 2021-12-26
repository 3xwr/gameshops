package repository

import (
	"gameservice/internal/model"
	"sync"
	"time"
)

type Cache struct {
	data    sync.Map
	timeout time.Duration
}

func NewCache(timeout time.Duration) *Cache {
	return &Cache{
		data:    sync.Map{},
		timeout: timeout,
	}
}

func (c *Cache) Load(key string) (model.GamePriceResponse, bool) {
	value, ok := c.data.Load(key)
	if !ok {
		return model.GamePriceResponse{}, false
	}
	p, ok := value.(model.GamePriceResponse)
	return p, ok
}

func (c *Cache) Store(key string, value model.GamePriceResponse) {
	t := time.Now()
	value.Timestamp = t
	c.data.Store(key, value)
}

func (c *Cache) GetTimeout() time.Duration {
	return c.timeout
}
