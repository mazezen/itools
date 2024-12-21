package luaredis

import (
	"errors"
	"github.com/redis/go-redis/v9"
)

type GoLuaRedisHash interface {
	HSetHashKey(rdb *redis.Client, key []string, field string, value interface{}) error
	HDelHashField(rdb *redis.Client, key []string, field string) error
	HGetHashFiled(rdb *redis.Client, key []string, field string) (interface{}, error)
	HExistsHashField(rdb *redis.Client, key []string, field string) (int, error)
	HGetAllHashField(rdb *redis.Client, key []string) ([]string, error)
	HIncrHashField(rdb *redis.Client, key []string, field string)
	HIncrByNumberHashField(rdb *redis.Client, key []string, field string, number int64)
	HIncrByFloat64HashField(rdb *redis.Client, key []string, field string, number float64)
	HIncrByFloat32HashField(rdb *redis.Client, key []string, field string, number float32)
	HKeysHash(rdb *redis.Client, key []string) ([]string, error)
	HValsHash(rdb *redis.Client, key []string) ([]interface{}, error)
	HLenHash(rdb *redis.Client, key []string) (int, error)
}

type GLRHash struct{}

func NewGLRHash() *GLRHash {
	return &GLRHash{}
}

// HSetHashKey set hash filed value
func (glh *GLRHash) HSetHashKey(rdb *redis.Client, key []string, field string, value interface{}) error {
	doCheckKey(key)
	doCheckHashField(field)
	err := setHashKey.Run(ctx, rdb, key, field, value).Err()
	return errIs(err)
}

// HDelHashField delete hash filed
func (glh *GLRHash) HDelHashField(rdb *redis.Client, key []string, field string) error {
	doCheckKey(key)
	doCheckHashField(field)
	err := delHashField.Run(ctx, rdb, key, field).Err()
	return errIs(err)
}

// HGetHashFiled get hash filed
func (glh *GLRHash) HGetHashFiled(rdb *redis.Client, key []string, field string) (interface{}, error) {
	doCheckKey(key)
	doCheckHashField(field)
	res, err := getHashFiled.Run(ctx, rdb, key, field).Result()
	if err = errIs(err); err != nil {
		return "", err
	}
	return res, nil
}

// HExistsHashField exists hash filed
func (glh *GLRHash) HExistsHashField(rdb *redis.Client, key []string, field string) (int, error) {
	doCheckKey(key)
	doCheckHashField(field)
	ret, err := hExistsHashField.Run(ctx, rdb, key, field).Int()
	if err != nil {
		return 0, err
	}
	return ret, nil
}

// HGetAllHashField get all hash field
func (glh *GLRHash) HGetAllHashField(rdb *redis.Client, key []string) ([]string, error) {
	doCheckKey(key)
	fields, err := hGetAllHashField.Run(ctx, rdb, key).StringSlice()
	if err != nil {
		return []string{}, err
	}
	return fields, nil
}

// HIncrHashField incr 1
func (glh *GLRHash) HIncrHashField(rdb *redis.Client, key []string, field string) {
	doCheckKey(key)
	doCheckHashField(field)
	hIncrHashField.Run(ctx, rdb, key, field, 1)
}

// HIncrByNumberHashField incr for given number
func (glh *GLRHash) HIncrByNumberHashField(rdb *redis.Client, key []string, field string, number int64) {
	doCheckKey(key)
	doCheckHashField(field)
	hIncrByNumberHashField.Run(ctx, rdb, key, field, number)
}

// HIncrByFloat64HashField incr for give float64 number
func (glh *GLRHash) HIncrByFloat64HashField(rdb *redis.Client, key []string, field string, number float64) {
	doCheckKey(key)
	doCheckHashField(field)
	hIncrByFloat64HashField.Run(ctx, rdb, key, field, number)
}

// HIncrByFloat32HashField incr for give float32 number
func (glh *GLRHash) HIncrByFloat32HashField(rdb *redis.Client, key []string, field string, number float32) {
	doCheckKey(key)
	doCheckHashField(field)
	hIncrByFloat32HashField.Run(ctx, rdb, key, field, number)
}

// HKeysHash return all keys list
func (glh *GLRHash) HKeysHash(rdb *redis.Client, key []string) ([]string, error) {
	doCheckKey(key)
	fields, err := hKeysHash.Run(ctx, rdb, key).StringSlice()
	if err != nil {
		return []string{}, err
	}
	return fields, nil
}

// HValsHash hash all valis
func (glh *GLRHash) HValsHash(rdb *redis.Client, key []string) ([]interface{}, error) {
	doCheckKey(key)
	return hValsHash.Run(ctx, rdb, key).Slice()
}

// HLenHash len hash
func (glh *GLRHash) HLenHash(rdb *redis.Client, key []string) (int, error) {
	doCheckKey(key)
	nums, err := hLenHash.Run(ctx, rdb, key).Int()
	if err != nil {
		return 0, err
	}
	return nums, nil
}

func errIs(err error) error {
	if errors.Is(err, redis.Nil) || errors.Is(err, nil) {
		return nil
	} else {
		return err
	}
}
