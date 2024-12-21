package luaredis

import (
	"github.com/redis/go-redis/v9"
)

var publish = redis.NewScript(`
	local channel = KEYS[1]
	local message = ARGV[1]
	return redis.call('PUBLISH', channel, message)
`)

var subscribe = redis.NewScript(`
	local channel = KEYS[1]
	local subscribe_client = redis.call('SUBSCRIBE', channel)
	
`)
