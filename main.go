package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i <= 8; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(10 * time.Second)
}
