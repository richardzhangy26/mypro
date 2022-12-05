package main

import (
	"fmt"
	"sync"
	"time"
)

var i int = 100

var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	defer wg.Done()
	lock.Lock()
	i += 1
	fmt.Printf("i++: %v\n", i)
	time.Sleep(time.Microsecond * 10)
	lock.Unlock()
}
func sub() {
	defer wg.Done()
	lock.Lock()
	i -= 1
	fmt.Printf("i--: %v\n", i)
	time.Sleep(time.Microsecond * 10)
	lock.Unlock()
}
func main() {
	for a := 0; a < 100; a++ {
		wg.Add(1)
		go add()
		wg.Add(1)
		go sub()
	}
	wg.Wait()
	fmt.Printf("end i: %v\n", i)
}
