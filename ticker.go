package itools

import "time"

func NewTicker(ttl time.Duration, sch chan string, fn func(sch chan string) int) {
	t := time.NewTicker(ttl)
	defer t.Stop()

	for range t.C {
		var res = fn(sch)
		switch res {
		case 1:
			t.Stop()
		default:
			t.Reset(ttl)
		}
	}
}
