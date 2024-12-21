package luaredis

import "github.com/redis/go-redis/v9"

type GoLuaRedisList interface {
	ListPush(rdb *redis.Client, key []string, value interface{}) int64
	ListLen(rdb *redis.Client, key []string) int64
	ListPop(rdb *redis.Client, key []string) interface{}
}

type GLRList struct{}

func NewGLRList() *GLRList {
	return &GLRList{}
}

// ListPush push element into list
func (lrl *GLRList) ListPush(rdb *redis.Client, key []string, value interface{}) int64 {
	doCheckKey(key)
	return listPush.Run(ctx, rdb, key, value).Val().(int64)
}

// ListLen list len
func (lrl *GLRList) ListLen(rdb *redis.Client, key []string) int64 {
	doCheckKey(key)
	return listLen.Run(ctx, rdb, key).Val().(int64)
}

// ListPop pop list
func (lrl *GLRList) ListPop(rdb *redis.Client, key []string) interface{} {
	doCheckKey(key)
	return listPop.Run(ctx, rdb, key).Val()
}
