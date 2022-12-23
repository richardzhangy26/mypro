package main

import (
	"fmt"
	"sync"
	"time"
)

var mymap = make(map[int]int, 10)
var lock sync.Mutex

func test(n int) {
	res := 1

	for i := 1; i <= n; i++ {
		res = res * i
	}
	lock.Lock()
	mymap[n] = res
	lock.Unlock()

}
func main() {
	for i := 1; i <= 20; i++ {
		go test(i)
	}
	time.Sleep(time.Second * 5)

	for k, v := range mymap {
		fmt.Printf("k:%v   v: %v\n", k, v)
	}

}
