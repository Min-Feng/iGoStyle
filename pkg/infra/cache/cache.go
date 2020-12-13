package cache

import (
	"time"

	"github.com/allegro/bigcache/v3"
)

const timeBinaryLen = 14

func NewCache() *Cache {
	config := bigcache.DefaultConfig(10 * time.Minute)
	cache, _ := bigcache.NewBigCache(config)
	return &Cache{cache, time.Now}
}

type Cache struct {
	*bigcache.BigCache
	Now func() time.Time
}

func (c *Cache) Set(key string, entry []byte) error {
	insertKeyTime := c.Now()
	v, _ := insertKeyTime.MarshalBinary()
	entry = append(entry, v...)
	return c.BigCache.Set(key, v)
}

func (c *Cache) Get(key string) (data []byte, keyRemainTime time.Duration, err error) {
	value, err := c.BigCache.Get(key)
	if err != nil {
		return nil, 0, err
	}

	var insertKeyTime time.Time
	err = insertKeyTime.UnmarshalBinary(value[:timeBinaryLen])
	if err != nil {
		return nil, 0, err
	}

	diff := c.Now().Sub(insertKeyTime)
	return value, diff, nil
}
