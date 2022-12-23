package main

import (
	"fmt"
)

func main() {
	var intChan chan int = make(chan int, 2)
	fmt.Printf("intChan: %v\nintChain address %p\n", intChan, &intChan)
	// 写入数据
	intChan <- 10
	num2 := <-intChan
	fmt.Printf("num2: %v\n", num2)
	num3 := <-intChan
	fmt.Printf("num3: %v\n", num3)
	close(intChan)
	num4 := <-intChan
	fmt.Printf("num4: %v\n", num4)
	num5 := <-intChan
	fmt.Printf("num5: %v\n", num5)

}
