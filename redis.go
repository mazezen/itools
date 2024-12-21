package itools

import (
	"context"
	"github.com/redis/go-redis/v9"
)

// NewRedis 初始化redis 连接
func NewRedis(addr, password string, db int) (*redis.Client, error) {
	rdc := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	_, err := rdc.Ping(context.TODO()).Result()
	if err != nil {
		return nil, err
	}
	return rdc, nil
}
