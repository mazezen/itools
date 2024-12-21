package luaredis

import (
	"github.com/redis/go-redis/v9"
)

var setSAdd = redis.NewScript(`
	local key = KEYS[1]
	local value = ARGV[1]
	return redis.call('SADD', key, value)
`)

var sMembers = redis.NewScript(`
	local key = KEYS[1]
	return redis.call('SMEMBERS', key)
`)
