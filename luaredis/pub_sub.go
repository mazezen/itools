package luaredis

import (
	"github.com/redis/go-redis/v9"
)

type GoLuaRedisPubSub interface {
	Publish(rdb *redis.Client, channel []string, message interface{}) (int, error)
	Subscribe(rdb *redis.Client, channel []string) interface{}
}

type GLRPubSub struct{}

func NewPubSub() *GLRPubSub {
	return &GLRPubSub{}
}

func (ps *GLRPubSub) Publish(rdb *redis.Client, channel []string, message interface{}) (int, error) {
	doCheckKey(channel)
	return publish.Run(ctx, rdb, channel, message).Int()
}

func (ps *GLRPubSub) Subscribe(rdb *redis.Client, channel []string) {
	doCheckKey(channel)
}
