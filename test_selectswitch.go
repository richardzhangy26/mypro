package main

import (
	"fmt"
	"time"
)

var chanint = make(chan int, 0)
var charstr = make(chan string)

func main() {
	go func() {
		chanint <- 100
		charstr <- "hello"
		defer close(chanint)
		defer close(charstr)
	}()
	for {
		select {
		case r := <-chanint:
			fmt.Printf("r: %v\n", r)

		case r := <-charstr:
			fmt.Printf("r: %v\n", r)
		default:
			fmt.Println("default..")
		}
		time.Sleep(time.Second)
	}
}
