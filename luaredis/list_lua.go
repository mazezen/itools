package luaredis

import (
	"github.com/redis/go-redis/v9"
)

var listPush = redis.NewScript(`
	local key = KEYS[1]
	local val = ARGV[1]
	return redis.call('LPUSH', key, val)
`)

var listLen = redis.NewScript(`
	local key = KEYS[1]
	return redis.call('LLEN', key)
`)

var listPop = redis.NewScript(`
	local key = KEYS[1]
	return redis.call('LPOP', key)
`)
