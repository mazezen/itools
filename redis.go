package itools

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
)

type RedisOption struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type RedisClientOption func(option *RedisOption)

func RedisClient(option ...RedisClientOption) *RedisOption {
	r := &RedisOption{}
	for _, o := range option {
		o(r)
	}
	return r
}

func WithAddress(address string) RedisClientOption {
	return func(o *RedisOption) {
		o.Addr = address
	}
}

func WithRedisPassword(password string) RedisClientOption {
	return func(o *RedisOption) {
		o.Password = password
	}
}

func WithRedisDB(db int) RedisClientOption {
	return func(o *RedisOption) {
		o.DB = db
	}
}

func (c *RedisOption) Connect() {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.DB,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		os.Exit(-1)
	}
	setClient(client)
}

var Rc *redis.Client

func setClient(_client *redis.Client) {
	Rc = _client
}
