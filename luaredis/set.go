package luaredis

import (
	"github.com/redis/go-redis/v9"
)

type GoLuaRedisSet interface {
	SetSAdd(rdb *redis.Client, key []string, value interface{}) int
	SMembers(rdb *redis.Client, key []string) ([]interface{}, error)
}

type GLRSet struct{}

func NewGLRSet() *GLRSet {
	return &GLRSet{}
}

// SetSAdd add element
func (lrs *GLRSet) SetSAdd(rdb *redis.Client, key []string, value interface{}) interface{} {
	doCheckKey(key)
	return setSAdd.Run(ctx, rdb, key, value).Val()
}

// SMembers members of set
func (lrs *GLRSet) SMembers(rdb *redis.Client, key []string) ([]interface{}, error) {
	doCheckKey(key)
	return sMembers.Run(ctx, rdb, key).Slice()
}
