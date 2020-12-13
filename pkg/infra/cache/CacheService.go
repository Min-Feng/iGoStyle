package cache

import (
	"errors"

	"github.com/allegro/bigcache/v3"
	"github.com/morikuni/failure"

	"AmazingTalker/pkg/technical/errutil"
)

type SaverFunc func(key string, data interface{}) error

func (fn SaverFunc) Save(key string, data interface{}) error {
	return fn(key, data)
}

func handleGetErr(err error) error {
	if errors.Is(err, bigcache.ErrEntryNotFound) {
		return failure.Translate(err, errutil.ErrNotFound)
	}
	return failure.Translate(err, errutil.ErrServer)
}
