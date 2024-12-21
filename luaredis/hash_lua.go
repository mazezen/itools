package luaredis

import (
	"github.com/redis/go-redis/v9"
)

var setHashKey = redis.NewScript(`
 local hashKey = KEYS[1]
 local field = ARGV[1]
 local value = ARGV[2]
 redis.call('HSET', hashKey, field, value)
`)

var delHashField = redis.NewScript(`
 local hashKey = KEYS[1]
 local field = ARGV[1]
 redis.call('HDEL', hashKey, field)
`)

var getHashFiled = redis.NewScript(`
 local hashKey = KEYS[1]
 local field = ARGV[1]
 return redis.call('HGET', hashKey, field)
`)

var hExistsHashField = redis.NewScript(`
 local hashKey = KEYS[1]
 local field = ARGV[1]
 return redis.call("HEXISTS", hashKey, field)
`)

var hGetAllHashField = redis.NewScript(`
 local hashKey = KEYS[1]
 return redis.call('HGETALL', hashKey)
`)

var hIncrHashField = redis.NewScript(`
 local hashKey = KEYS[1]
 local field = ARGV[1]
 redis.call("HINCRBY", hashKey, field, 1)
`)

var hIncrByNumberHashField = redis.NewScript(`
 local hashKey = KEYS[1]
 local field = ARGV[1]
 local number = ARGV[2]
 redis.call('HINCRBY', hashKey, field, number)
`)

var hIncrByFloat64HashField = redis.NewScript(`
 local hashKey = KEYS[1]
 local field = ARGV[1]
 local number = ARGV[2]
 redis.call('HINCRBYFLOAT', hashKey, field, number)
`)

var hIncrByFloat32HashField = redis.NewScript(`
 local hashKey = KEYS[1]
 local field = ARGV[1]
 local number = ARGV[2]
 redis.call('HINCRBYFLOAT', hashKey, field, number)
`)

var hKeysHash = redis.NewScript(`
 local hashKey = KEYS[1]
 return redis.call('HKEYS', hashKey)
`)

var hLenHash = redis.NewScript(`
 local hashKey = KEYS[1]
 return redis.call('HLEN', hashKey)
`)

var hValsHash = redis.NewScript(`
 local hashKey = KEYS[1]
 return redis.call('HVALS', hashKey)
`)
