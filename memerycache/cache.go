package memerycache

import (
	"fmt"
	"sync"
	"time"
)

type Cache interface {
	SetMaxMemory(size string) bool
	Set(key string, val interface{}, expire time.Duration) bool
	Get(key string) (interface{}, bool)
	Del(key string) bool
	Exists(key string) bool
	Flush() bool
	Keys() int64
}

type MemoryCache struct {
	maxMemorySize           int64
	maxMemorySizeS          string
	currentMemorySize       int64
	cMap                    map[string]*cMapValue
	mLock                   sync.RWMutex  // 读写锁
	clearExpireKeysInterval time.Duration // 清除过期key时间周期
}

type cMapValue struct {
	value   interface{}   // value值
	expireT time.Time     // 过期时间
	expire  time.Duration // 有效时长
	size    int64         // value大小
}

func NewMemCache() Cache {
	m := &MemoryCache{
		cMap:                    make(map[string]*cMapValue),
		clearExpireKeysInterval: time.Second * 10,
	}
	go m.clearExpireKeys()
	return m
}

var defaultMemorySize int64 = 100 // 100MB

// SetMaxMemory size: 1KB 100KB 1MB 100MB 1GB default 100MB
func (m *MemoryCache) SetMaxMemory(size string) bool {
	m.maxMemorySize, m.maxMemorySizeS = CovertSize(size, defaultMemorySize)
	fmt.Println(m.maxMemorySize, m.maxMemorySizeS)
	return false
}

func (m *MemoryCache) Set(key string, val interface{}, expire time.Duration) bool {
	m.mLock.Lock()
	defer m.mLock.Unlock()
	v := &cMapValue{
		value:   val,
		expireT: time.Now().Add(expire),
		expire:  expire,
		size:    CalculateSize(val),
	}
	m.del(key)
	m.add(key, v)
	if m.currentMemorySize > m.maxMemorySize {
		m.del(key)
		panic("max memory size is not enough")
	}
	return true
}

func (m *MemoryCache) get(key string) (*cMapValue, bool) {
	val, ok := m.cMap[key]
	return val, ok
}

func (m *MemoryCache) del(key string) {
	tmp, ok := m.get(key)
	if ok && tmp != nil {
		m.currentMemorySize -= tmp.size
		delete(m.cMap, key)
	}
}

func (m *MemoryCache) add(key string, val *cMapValue) {
	m.cMap[key] = val
	m.currentMemorySize += val.size
}

func (m *MemoryCache) Get(key string) (interface{}, bool) {
	m.mLock.RLock()
	defer m.mLock.RUnlock()
	v, ok := m.get(key)
	if ok {
		if v.expire != 0 && v.expireT.Before(time.Now()) {
			m.del(key)
			return nil, false
		}
	}
	return v.value, ok
}

func (m *MemoryCache) Del(key string) bool {
	m.mLock.Lock()
	defer m.mLock.Unlock()
	m.del(key)
	return true
}

func (m *MemoryCache) Exists(key string) bool {
	m.mLock.RLock()
	defer m.mLock.RUnlock()
	_, ok := m.cMap[key]
	return ok
}

func (m *MemoryCache) Flush() bool {
	m.mLock.Lock()
	defer m.mLock.Unlock()
	m.cMap = make(map[string]*cMapValue, 0)
	m.currentMemorySize = 0
	return true
}

func (m *MemoryCache) Keys() int64 {
	m.mLock.RLock()
	defer m.mLock.RUnlock()
	return int64(len(m.cMap))
}

func (m *MemoryCache) clearExpireKeys() {
	tk := time.NewTicker(m.clearExpireKeysInterval)
	defer tk.Stop()
	for {
		select {
		case <-tk.C:
			for k, item := range m.cMap {
				if item.expire != 0 && time.Now().After(item.expireT) {
					m.mLock.Lock()
					m.del(k)
					m.mLock.Unlock()
				}
			}
		}
	}
}
