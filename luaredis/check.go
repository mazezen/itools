package luaredis

import (
	"errors"
	"reflect"
)

var (
	KeyIsEmpty       = errors.New("key is allowed not empty")
	ValuesIsEmpty    = errors.New("value is allowed not empty")
	TtlIsNoAllowed   = errors.New("ttl is not allowed 0 or less than 0")
	DbIsNoAllowed    = errors.New("ttl is not less than 0")
	HashFieldIsEmpty = errors.New("hash field is empty")
)

// doCheckKey
func doCheckKey(key []string) {
	if len(key) == 0 {
		panic(KeyIsEmpty)
	}
}

// doCheckTtl
func doCheckTtl(ttl int) {
	if ttl <= 0 {
		panic(TtlIsNoAllowed)
	}
}

// doCheckValue
func doCheckValue(value interface{}) {
	if isEmpty(value) {
		panic(ValuesIsEmpty)
	}
}

func isEmpty(v interface{}) bool {
	if v == nil {
		return true
	}
	value := reflect.ValueOf(v)
	switch value.Kind() {
	case reflect.Ptr, reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return value.IsNil() || value.Len() == 0
	case reflect.Struct:
		zero := reflect.Zero(value.Type())
		return reflect.DeepEqual(value.Interface(), zero.Interface())
	}
	return false
}

// doCheckKeyTtl
func doCheckKeyTtl(key []string, ttl int) {
	if len(key) == 0 {
		panic(errors.New("key is allowed not empty"))
	}

	if ttl <= 0 {
		panic(TtlIsNoAllowed)
	}
}

// doCheckDb
func doCheckDb(db int) {
	if db < 0 {
		panic(DbIsNoAllowed)
	}
}

// doCheckHashField
func doCheckHashField(field string) {
	if field == "" {
		panic(HashFieldIsEmpty)
	}
}
