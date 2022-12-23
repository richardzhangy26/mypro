package main

import (
	"fmt"
)

func writeData(intChain chan int) {
	for i := 0; i < 50; i++ {
		intChain <- i
		fmt.Printf("write data %v\n", i)
	}
	close(intChain)
}
func readData(intChain chan int, exitChain chan bool) {
	for {
		v, ok := <-intChain
		if !ok {
			break
		}
		fmt.Printf("read data : %v\n", v)
	}
	// 读取完成后，任务完成
	exitChain <- true
	close(exitChain)
}
func main() {

	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
