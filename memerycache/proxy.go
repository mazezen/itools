package memerycache

import "time"

type CacheImpl struct {
	memCache Cache
}

func NewCache() *CacheImpl {
	return &CacheImpl{memCache: NewMemCache()}
}

func (c *CacheImpl) SetMaxMemory(size string) bool {
	return c.memCache.SetMaxMemory(size)
}

func (c *CacheImpl) Set(key string, val interface{}, expire ...time.Duration) bool {
	expireT := time.Second * 0
	if len(expire) > 0 {
		expireT = expire[0]
	}
	return c.memCache.Set(key, val, expireT)
}

func (c *CacheImpl) Get(key string) (interface{}, bool) {
	return c.memCache.Get(key)
}

func (c *CacheImpl) Del(key string) bool {
	return c.memCache.Del(key)
}

func (c *CacheImpl) Exists(key string) bool {
	return c.memCache.Exists(key)
}

func (c *CacheImpl) Flush() bool {
	return c.memCache.Flush()
}

func (c *CacheImpl) Keys() int64 {
	return c.memCache.Keys()
}
