package itools

import (
	"encoding/json"
	"sync"
	"time"

	lru "github.com/hashicorp/golang-lru"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/util"
)

var (
	ErrEmptyKey = errors.New("key could not be empty")
)

type Database interface {
	Put(key string, value interface{}) error
	BatchPut(kvs map[string]interface{}) error
	Get(key string) ([]byte, error)
	Has(key string) (bool, error)
	Delete(key string) error
	BatchDelete(keys []string) error
	SelectAll() iterator.Iterator
	SelectPrefixSubsetKeyAll(key string) ([]map[string]interface{}, error)
	CountPrefixSubsetKey(key string) (int64, error)
	CountAll() (int64, error)
	DeletePrefixSubsetKey(key string) (bool, error)
}

type LevelDB struct {
	db *leveldb.DB
}

type cacheEntry struct {
	value     []byte
	timestamp time.Time
}

type LevelDBWithCache struct {
	db    *LevelDB
	cache *lru.Cache
	mutex sync.Mutex
}

const (
	cacheSize = 100000        // LRU 缓存大小
	cacheTTL  = time.Hour   // 缓存过期时间为1小时
)

func NewLevelDB(path string) (*LevelDB, error) {
	db, err := leveldb.OpenFile(path, nil)
	if _, corrupted := err.(*errors.ErrCorrupted); corrupted {
		db, err = leveldb.RecoverFile(path, nil)
	}
	if err != nil {
		return nil, err
	}
	result := &LevelDB{
		db: db,
	}
	return result, nil
}

func NewLevelDBWithCache(path string) (*LevelDBWithCache, error) {
	db, err := NewLevelDB(path)
	if err != nil {
		return nil, err
	}
	cache, err := lru.New(cacheSize)
	if err != nil {
		return nil, err
	}
	return &LevelDBWithCache{
		db:    db,
		cache: cache,
	}, nil
}

func (db *LevelDBWithCache) Put(key string, value interface{}) error {
	if len(key) < 1 {
		return ErrEmptyKey
	}
	res, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = db.db.db.Put([]byte(key), res, nil)
	if err == nil {
		db.mutex.Lock()
		defer db.mutex.Unlock()
		db.cache.Add(key, cacheEntry{
			value:     res,
			timestamp: time.Now(),
		})
	}
	return err
}

func (db *LevelDBWithCache) BatchPut(kvs map[string]interface{}) error {
	if len(kvs) == 0 {
		return nil
	}
	batch := new(leveldb.Batch)
	db.mutex.Lock()
	defer db.mutex.Unlock()
	for key, value := range kvs {
		if len(key) < 1 {
			return ErrEmptyKey
		}
		res, err := json.Marshal(value)
		if err != nil {
			return err
		}
		batch.Put([]byte(key), res)
		db.cache.Add(key, cacheEntry{
			value:     res,
			timestamp: time.Now(),
		})
	}
	return db.db.db.Write(batch, nil)
}

func (db *LevelDBWithCache) Get(key string) ([]byte, error) {
	if len(key) < 1 {
		return nil, ErrEmptyKey
	}
	db.mutex.Lock()
	if val, ok := db.cache.Get(key); ok {
		if entry, ok := val.(cacheEntry); ok {
			if time.Since(entry.timestamp) <= cacheTTL {
				db.mutex.Unlock()
				return entry.value, nil
			}
			db.cache.Remove(key)
		}
	}
	db.mutex.Unlock()

	data, err := db.db.db.Get([]byte(key), nil)
	if err == nil {
		db.mutex.Lock()
		defer db.mutex.Unlock()
		db.cache.Add(key, cacheEntry{
			value:     data,
			timestamp: time.Now(),
		})
	}
	return data, err
}

func (db *LevelDBWithCache) Has(key string) (bool, error) {
	if len(key) < 1 {
		return false, ErrEmptyKey
	}
	db.mutex.Lock()
	if val, ok := db.cache.Get(key); ok {
		if entry, ok := val.(cacheEntry); ok {
			if time.Since(entry.timestamp) <= cacheTTL {
				db.mutex.Unlock()
				return true, nil
			}
			db.cache.Remove(key)
		}
	}
	db.mutex.Unlock()

	exists, err := db.db.db.Has([]byte(key), nil)
	if err == nil && exists {
		db.mutex.Lock()
		defer db.mutex.Unlock()
		db.cache.Add(key, cacheEntry{
			value:     []byte{},
			timestamp: time.Now(),
		})
	}
	return exists, err
}

func (db *LevelDBWithCache) Delete(key string) error {
	if len(key) < 1 {
		return ErrEmptyKey
	}
	err := db.db.db.Delete([]byte(key), nil)
	if err == nil {
		db.mutex.Lock()
		defer db.mutex.Unlock()
		db.cache.Remove(key)
	}
	return err
}

func (db *LevelDBWithCache) BatchDelete(keys []string) error {
	if len(keys) == 0 {
		return nil
	}
	batch := new(leveldb.Batch)
	db.mutex.Lock()
	defer db.mutex.Unlock()
	for _, key := range keys {
		if len(key) < 1 {
			return ErrEmptyKey
		}
		batch.Delete([]byte(key))
		db.cache.Remove(key)
	}
	return db.db.db.Write(batch, nil)
}

func (db *LevelDBWithCache) DeleteAll() (bool, error) {
	iter := db.db.db.NewIterator(nil, nil)
	batch := new(leveldb.Batch)
	db.mutex.Lock()
	defer db.mutex.Unlock()
	for iter.Next() {
		batch.Delete(iter.Key())
		db.cache.Remove(string(iter.Key()))
	}
	iter.Release()
	if err := iter.Error(); err != nil {
		return false, err
	}
	return true, db.db.db.Write(batch, nil)
}

func (db *LevelDBWithCache) DeletePrefixSubsetKey(key string) (bool, error) {
	iter := db.db.db.NewIterator(util.BytesPrefix([]byte(key)), nil)
	batch := new(leveldb.Batch)
	db.mutex.Lock()
	defer db.mutex.Unlock()
	for iter.Next() {
		batch.Delete(iter.Key())
		db.cache.Remove(string(iter.Key()))
	}
	iter.Release()
	if err := iter.Error(); err != nil {
		return false, err
	}
	return true, db.db.db.Write(batch, nil)
}

func (db *LevelDBWithCache) SelectAll() iterator.Iterator {
	return db.db.db.NewIterator(nil, nil)
}

func (db *LevelDBWithCache) SelectPrefixSubsetKeyAll(key string) ([]map[string]interface{}, error) {
	m := make(map[string]interface{})
	ms := make([]map[string]interface{}, 0)

	iter := db.db.db.NewIterator(util.BytesPrefix([]byte(key)), nil)
	for iter.Next() {
		k := string(iter.Key())
		v := string(iter.Value())
		m[k] = v
		ms = append(ms, m)
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		return nil, err
	}
	return ms, nil
}

func (db *LevelDBWithCache) CountPrefixSubsetKey(key string) (int64, error) {
	var sum int64 = 0
	iter := db.db.db.NewIterator(util.BytesPrefix([]byte(key)), nil)
	for iter.Next() {
		sum++
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		return 0, err
	}
	return sum, nil
}

func (db *LevelDBWithCache) CountAll() (int64, error) {
	var sum int64 = 0
	iter := db.db.db.NewIterator(nil, nil)
	for iter.Next() {
		sum++
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		return 0, err
	}
	return sum, nil
}