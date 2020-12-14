package cache

import (
	"errors"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/morikuni/failure"

	"iGoStyle/pkg/technical/errutil"
)

const timeBinaryLen = 15

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
	// log.Debug().Int("time byte len", len(v)).Send()
	v = append(v, entry...)
	return failure.Translate(c.BigCache.Set(key, v), errutil.ErrServer)
}

func (c *Cache) Get(key string) (data []byte, keyRemainTime time.Duration, err error) {
	value, err := c.BigCache.Get(key)
	if err != nil {
		if errors.Is(err, bigcache.ErrEntryNotFound) {
			return nil, 0, failure.Translate(err, errutil.ErrNotFound)
		}
		return nil, 0, failure.Translate(err, errutil.ErrServer)
	}

	var insertKeyTime time.Time
	err = insertKeyTime.UnmarshalBinary(value[:timeBinaryLen])
	if err != nil {
		return nil, 0, failure.Wrap(err)
	}

	diff := c.Now().Sub(insertKeyTime)
	return value[timeBinaryLen:], diff, nil
}
