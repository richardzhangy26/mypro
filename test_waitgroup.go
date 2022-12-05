package main

import (
	"fmt"
	"sync"
)

var wp sync.WaitGroup

func showMsg(i int) {
	defer wp.Done()
	fmt.Printf("i: %v\n", i)
}
func main() {
	for i := 0; i < 10; i++ {
		go showMsg(i)
		wp.Add(1)
	}
	wp.Wait()
	fmt.Println("end...")
}
