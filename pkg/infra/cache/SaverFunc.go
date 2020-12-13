package cache

type SaverFunc func(key string, data interface{}) error

func (fn SaverFunc) Save(key string, data interface{}) error {
	return fn(key, data)
}
