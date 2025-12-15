package itools

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomPollingNext(t *testing.T) {
	rp := &RandomPolling{}
	err := rp.Add("127.0.0.1:2003", "127.0.0.1:2004")
	assert.NotErrorIs(t, err, ParamsNoEnoughError)
	n := rp.RandomNext()
	t.Logf("random polling: %#v", n)

	rap := &RotationPolling{}
	err = rap.Add("127.0.0.1:2003", "127.0.0.1:2004")
	assert.NotErrorIs(t, err, ParamsNoEnoughError)
	rotationNext1 := rap.RotationNext()
	rotationNext2 := rap.RotationNext()
	t.Logf("rotation polling: %#v", rotationNext1)
	t.Logf("rotation polling: %#v", rotationNext2)

	hashPolling := NewConsistentHashPolling(10, nil)
	hashPolling.Add("127.0.0.1:2003")
	hashPolling.Add("127.0.0.1:2004")
	hashPolling.Add("127.0.0.1:2005")
	hashPolling.Add("127.0.0.1:2006")
	hashPolling.Add("127.0.0.1:2007")

	fmt.Println(hashPolling.Next("http://127.0.0.1:2002/base/getinfo"))
	fmt.Println(hashPolling.Next("http://127.0.0.1:2002/base/errinfo"))
	fmt.Println(hashPolling.Next("http://127.0.0.1:2002/base/errinfo"))
	fmt.Println(hashPolling.Next("http://127.0.0.1:2002/base/pwd"))

	fmt.Println(hashPolling.Next("127.0.0.1"))
	fmt.Println(hashPolling.Next("192.168.0.1"))
	fmt.Println(hashPolling.Next("127.0.0.1"))

	r := &WeightPolling{}
	r.Add("127.0.0.1:2003", "5")
	r.Add("127.0.0.1:2004", "3")
	r.Add("127.0.0.1:2005", "2")

	fmt.Println(r.Next())
	fmt.Println(r.Next())
	fmt.Println(r.Next())
	fmt.Println(r.Next())
	fmt.Println(r.Next())
	fmt.Println(r.Next())
	fmt.Println(r.Next())
	fmt.Println(r.Next())
	fmt.Println(r.Next())
	fmt.Println(r.Next())
}
