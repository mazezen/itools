package main

import (
	"github.com/mazezen/itools"
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var lc itools.Counter
	lc.Set(10, time.Second) // 1s内10速率
	for i := 0; i < 50; i++ {
		wg.Add(1)
		log.Println("创建请求: ", i)
		go func(i int) {

			if lc.Pass() {
				log.Println("response: ", i)
			}
			wg.Done()
		}(i)
		time.Sleep(10 * time.Millisecond)
	}
	wg.Wait()
}
