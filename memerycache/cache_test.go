package memerycache

import (
	"fmt"
	"testing"
	"time"
)

func TestCacheOperate(t *testing.T) {
	d := []struct {
		key    string
		value  interface{}
		expire time.Duration
	}{
		{"fsd", 2323, time.Second * 11},
		{"fsd1", false, time.Second * 22},
		{"fsd2", true, time.Second * 33},
		{"fsd3", "fsdfs", time.Second * 44},
		{"fsd4", "21321", time.Second * 55},
		{"fsd5", "dssdfdsfs", time.Second * 66},
		{"fsd6", 122, time.Second * 77},
		{"fsd7", 23, time.Second * 88},
		{"fsd8", 2, time.Second * 98},
		{"fsd9", 3, time.Second * 88},
		{"fsd10", map[string]interface{}{"a": "100", "b": false}, time.Second * 13},
		{"fsd11", 89, time.Second * 45},
		{"fsd12", 223, time.Second * 67},
		{"fsd13", 322, time.Second * 89},
		{"fsd14", 454, time.Second * 90},
		{"fsd15", 878, time.Second * 56},
		{"fsd16", "fdfsd", time.Second * 78},
	}
	c := NewCache()
	c.SetMaxMemory("10MB")
	for _, item := range d {
		c.Set(item.key, item.value, item.expire)
		val, ok := c.Get(item.key)
		if !ok {
			t.Error("获取失败")
		}
		if item.key != "fsd10" && val != item.value {
			t.Error("获取缓存的值不一致")
		}
		_, ok1 := val.(map[string]interface{})
		if item.key == "fsd10" && !ok1 {
			t.Error("获取缓存的值不一致")
		}
	}
	if int64(len(d)) != c.Keys() {
		t.Error("缓存的数量不一致")
	}
	c.Del(d[0].key)
	c.Del(d[0].key)
	if int64(len(d)) != c.Keys()+2 {
		t.Error("缓存删除失败")
	}
	time.Sleep(time.Second * 100)

	if c.Keys() != 0 {
		t.Error("缓存清空失败")
	}
	println("测试成功...")
}

func TestMemoryCache_Set(t *testing.T) {
	c := NewCache()
	c.SetMaxMemory("10MB")
	c.Set("a", 10, time.Second*50)
	println("set success...")
}

func TestMemoryCache_Del(t *testing.T) {
	c := NewCache()
	ok := c.Del("a")
	fmt.Println(ok)
}

func TestMemoryCache_Get(t *testing.T) {
	c := NewCache()
	v, ok := c.Get("a")
	fmt.Println(v)
	fmt.Println(ok)
}
