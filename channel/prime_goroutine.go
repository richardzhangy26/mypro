// 便求统计1-200000 的数字中，哪些是素数？这个问题在本章开篇就提出了，现在我们有 goroutine 和 channel 的知识后，就可以完成了[测试数据：80000]
package main

import (
	"fmt"
)

func putNum(intChan chan int) {
	for i := 1; i <= 11000; i++ {
		intChan <- i

	}
	close(intChan)
}
func primeNum(primeChan chan int, intChan chan int, exitChan chan bool) {
	var flag bool
	// exitChan <- true
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		for i := 2; i < num/2; i++ {
			if num%i == 0 {
				flag = false
				break
				fmt.Printf("flag: %v\n", flag)
			}
		}
		if flag {
			primeChan <- num
		}
		// fmt.Printf("flag: %v\n", flag)
	}
	fmt.Println("有一个协程取不到数据，所以退出")
	exitChan <- true
	// close(exitChan)
}
func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 9000)
	exitChan := make(chan bool, 4)
	go putNum(intChan)
	// 开启4个协程，从goroutine取出数据，计算是否为素数

	for i := 0; i < 4; i++ {
		go primeNum(primeChan, intChan, exitChan)
	}
	// time.Sleep(time.Second * 2)
	// go func() {
	// 	for i := 0; i < 4; i++ {
	// 		<-exitChan
	// 	}
	// 	close(primeChan)
	// }()
	// time.Sleep(time.Second * 4)

	for i := 0; i < 4; i++ {
		// time.Sleep(time.Second * 1)
		num := <-exitChan
		fmt.Printf("num: %v\n", num)
	}

	close(primeChan)
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("res: %v\n", res)
	}
	fmt.Println("主线程结束")
	// count := 0
	// for {
	// 	flag, ok := <-exitChan
	// 	if flag {
	// 		count++
	// 	}
	// }

}
