package luaredis

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
)

type GoLuaRedisString interface {
	SetWithExpire(rdb *redis.Client, key []string, val string, ttl int) (string, error)
	SetWithoutExpire(rdb *redis.Client, key []string, val string) (string, error)
	GetKey(rdb *redis.Client, key []string) (string, error)
	GetOrSetWithExpire(rdb *redis.Client, key []string, val string, ttl int) (string, error)
	GetOrSetWithoutExpire(rdb *redis.Client, key []string, val string) (string, error)
	DelKey(rdb *redis.Client, key []string) (int64, error)
	ExistsKey(rdb *redis.Client, key []string) (int64, error)
	ExpireKey(rdb *redis.Client, key []string, ttl int) (int64, error)
	PatternKeys(rdb *redis.Client, key []string) ([]string, error)
	PtlKey(rdb *redis.Client, key []string) (int64, error)
	MoveKeyToDb(rdb *redis.Client, key []string, db int) (int64, error)
}

type GLRedis struct{}

func NewGLRedis() *GLRedis {
	return &GLRedis{}
}

var ctx = context.Background()

// SetWithExpire set value with expire for given key
// unit time second
func (glr *GLRedis) SetWithExpire(rdb *redis.Client, key []string, val string, ttl int) (string, error) {
	doCheckKeyTtl(key, ttl)
	ret, err := setWithExpire.Run(ctx, rdb, key, val, ttl).Result()
	if err != nil {
		return "", err
	}
	return ret.(string), nil
}

// SetWithoutExpire set value without expire for given key
func (glr *GLRedis) SetWithoutExpire(rdb *redis.Client, key []string, val string) (string, error) {
	doCheckKey(key)
	ret, err := setWithoutExpire.Run(ctx, rdb, key, val).Result()
	if err != nil {
		return "", err
	}
	return ret.(string), nil
}

// GetKey get value for given key
func (glr *GLRedis) GetKey(rdb *redis.Client, key []string) (string, error) {
	doCheckKey(key)
	ret, err := getKey.Run(ctx, rdb, key).Result()
	if err != nil {
		return "", err
	}
	return ret.(string), nil
}

// GetOrSetWithExpire get not exist, do set
// unit time second
func (glr *GLRedis) GetOrSetWithExpire(rdb *redis.Client, key []string, val string, ttl int) (string, error) {
	doCheckKeyTtl(key, ttl)
	ret, err := getOrSetWithExpire.Run(ctx, rdb, key, val, ttl).Result()
	if err != nil {
		return "", err
	}
	return ret.(string), nil
}

// GetOrSetWithoutExpire get not exist, do set
func (glr *GLRedis) GetOrSetWithoutExpire(rdb *redis.Client, key []string, val string) (string, error) {
	doCheckKey(key)
	ret, err := getOrSetWithoutExpire.Run(ctx, rdb, key, val).Result()
	if err != nil {
		return "", err
	}
	return ret.(string), nil
}

// DelKey del for given key
// success return 1, failed return 0
func (glr *GLRedis) DelKey(rdb *redis.Client, key []string) (int64, error) {
	doCheckKey(key)
	ret, err := delKey.Run(ctx, rdb, key).Result()
	if err != nil {
		return ret.(int64), err
	}
	return ret.(int64), nil
}

// ExistsKey check key is exists for given key
// exists return 1, is not exists return 0
func (glr *GLRedis) ExistsKey(rdb *redis.Client, key []string) (int64, error) {
	doCheckKey(key)
	ret, err := existsKey.Run(ctx, rdb, key).Result()
	if err != nil {
		return ret.(int64), err
	}
	return ret.(int64), err
}

// ExpireKey
// success return 1, failed return 0
// unit time second
func (glr *GLRedis) ExpireKey(rdb *redis.Client, key []string, ttl int) (int64, error) {
	doCheckKey(key)
	doCheckTtl(ttl)
	ret, err := expireKey.Run(ctx, rdb, key, ttl).Result()
	if err != nil {
		return ret.(int64), err
	}
	return ret.(int64), nil
}

// PatternKeys return list for given pattern keys
func (glr *GLRedis) PatternKeys(rdb *redis.Client, key []string) ([]string, error) {
	doCheckKey(key)
	ret, err := patternKeys.Run(ctx, rdb, key).StringSlice()
	if err != nil {
		return nil, err
	}

	if len(ret) == 0 {
		return nil, errors.New("no pattern keys")
	}

	return ret, nil
}

// PtlKey return ttl for given key
// return unit time second
func (glr *GLRedis) PtlKey(rdb *redis.Client, key []string) (int64, error) {
	doCheckKey(key)
	ret, err := ptlKey.Run(ctx, rdb, key).Result()
	if err != nil {
		return ret.(int64), err
	}
	return ret.(int64) / 1e3, nil
}

// MoveKeyToDb move to selected db for given key
func (glr *GLRedis) MoveKeyToDb(rdb *redis.Client, key []string, db int) (int64, error) {
	doCheckKey(key)
	doCheckDb(db)
	ret, err := moveKeyToDb.Run(ctx, rdb, key, db).Result()
	if err != nil {
		return ret.(int64), err
	}
	return ret.(int64), nil
}
