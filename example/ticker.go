package main

import (
	"fmt"
	"github.com/mazezen/itools"
	"log"
	"net/http"
	"strconv"
	"time"
)

func doSomething(sch chan string) int {
	selected, ok := <-sch
	if !ok && selected == "" {
		return -1
	}
	fmt.Println("selected is", selected)

	// do something
	if selected == "1" {
		fmt.Println("select is 1 and stop ticker")
		return 1
	}
	fmt.Println("select is not 1 and go ticker")
	sch <- selected
	return -1
}

func main() {

	http.Handle("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	}))

	http.Handle("/test", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//i := 1
		i := 2
		selected := strconv.Itoa(i)
		var sch = make(chan string, 1)
		sch <- selected

		// 创建定时器
		go time.AfterFunc(0, func() {
			itools.NewTicker(time.Second*2, sch, doSomething)
		})

	}))

	log.Fatal(http.ListenAndServe(":8080", nil))

}
