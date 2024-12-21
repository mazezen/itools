package luaredis

import (
	"github.com/redis/go-redis/v9"
)

var setWithExpire = redis.NewScript(`
	local key = KEYS[1]
    local value = ARGV[1]
    local ttl = ARGV[2]
    redis.call('SET', key, value)
    redis.call('EXPIRE', key, ttl)
    return value
`)

var setWithoutExpire = redis.NewScript(`
	local key = KEYS[1]
    local value = ARGV[1]
    redis.call('SET', key, value)
    return value
`)

var getKey = redis.NewScript(`
	local key = KEYS[1]
    local result = redis.call('GET', key)
    if result then
        return result
    end
`)

var getOrSetWithExpire = redis.NewScript(`
	local key = KEYS[1]
    local value = ARGV[1]
    local ttl = ARGV[2]
    local result = redis.call('GET', key)
    if result then
        return result
    end
    redis.call('SET', key, value)
    redis.call('EXPIRE', key, ttl)
    return value
`)

var getOrSetWithoutExpire = redis.NewScript(`
	local key = KEYS[1]
 	local value = ARGV[1]
	local result = redis.call('GET', key)
    if result then
        return result
    end
	redis.call('SET', key, value)
    return value
`)

var delKey = redis.NewScript(`
	local key = KEYS[1]
	local result = redis.call('DEL', key)
	if result then
		return result
	end
`)

var existsKey = redis.NewScript(`
	local key = KEYS[1]
	local result = redis.call('EXISTS', key)
	if result then
		return result
	end
`)

var expireKey = redis.NewScript(`
	local key = KEYS[1]
	local ttl = ARGV[1]
	local result = redis.call('EXPIRE', key, ttl)
	if result then
		return result
	end
`)

var patternKeys = redis.NewScript(`
	local key = KEYS[1]
	return redis.call('KEYS', key)
`)

var ptlKey = redis.NewScript(`
	local key = KEYS[1]
	return redis.call('PTTL', key)
`)

var moveKeyToDb = redis.NewScript(`
	local key = KEYS[1]
	local db = ARGV[1]
	return redis.call('MOVE', key, db)
`)
