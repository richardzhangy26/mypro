package main

import (
	"fmt"
)

var c = make(chan int)

func main() {
	go func() {
		for i := 0; i < 2; i++ {
			c <- i
			// time.Sleep(time.Second * 2)
		}
		// close(c)
	}()
	r := <-c
	fmt.Printf("r: %v\n", r)
	r = <-c
	fmt.Printf("r: %v\n", r)
	r = <-c
	fmt.Printf("r: %v\n", r)
	for v := range c {
		fmt.Printf("v: %v\n", v)
	}
	// for ok;v :=range c{
	// 	if ok ==nil {

	// 	}
	// 	else {

	// 	}

	// }
	fmt.Println("end...")
}
