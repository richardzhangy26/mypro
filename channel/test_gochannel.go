package main

import (
	"fmt"
)

func main() {
	// exitChan := make(chan bool, 1)

	// go func() {
	// 	exitChan <- true
	// 	fmt.Printf("go i: %v\n", 1)
	// }()

	// go func() {
	// 	exitChan <- true
	// 	fmt.Printf("go i: %v\n", 2)
	// }()
	// go func() {
	// 	exitChan <- true
	// 	fmt.Printf("go i: %v\n", 3)
	// }()
	// go func() {
	// 	for i := 0; i < 3; i++ {
	// 		// time.Sleep(time.Second * 5)
	// 		num := <-exitChan
	// 		fmt.Printf("num: %v\n", num)
	// 	}
	// }()
	// time.Sleep(time.Second * 2)
	for i := 0; i < 1000; i++ {
		go func() {
			fmt.Println("the goroutine", i)
		}()
	}
}
