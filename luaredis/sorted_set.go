package luaredis

import (
	"errors"
	"github.com/redis/go-redis/v9"
)

type GoLuaRedisSortedSet interface {
	SetZAdd(rdb *redis.Client, key []string, score interface{}, value interface{}) error
	SetZRange(rdb *redis.Client, key []string, start, end int) ([]interface{}, error)
}

type GLRSortedSet struct{}

func NewGLRSortedSet() *GLRSortedSet {
	return &GLRSortedSet{}
}

// SetZAdd add sorted set element
func (lss *GLRSortedSet) SetZAdd(rdb *redis.Client, key []string, score, value interface{}) error {
	doCheckKey(key)
	return setZAdd.Run(ctx, rdb, key, score, value).Err()
}

// SetZRange range of sorted
func (lss *GLRSortedSet) SetZRange(rdb *redis.Client, key []string, start, end int) ([]interface{}, error) {
	doCheckKey(key)
	zSet, err := setZRange.Run(ctx, rdb, key, start, end).Slice()
	if errors.Is(err, redis.Nil) || err == nil {
		return zSet, nil
	} else {
		return nil, err
	}
}
