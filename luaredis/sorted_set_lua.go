package luaredis

import (
	"github.com/redis/go-redis/v9"
)

var setZAdd = redis.NewScript(`
	local key = KEYS[1]
	local score = ARGV[1]
	local val = ARGV[2]
	redis.call("ZADD", key, score, val)
`)

var setZRange = redis.NewScript(`
	local key = KEYS[1]
	local range_start = ARGV[1]
	local range_end = ARGV[2]
	return redis.call("ZRANGE", key, range_start, range_end)
`)
