// 1) 启动一个协程，将1-2000的数放入到一个channel中，比如 numChan
// 2)启动8个物程、从numChan取出数(比如n)，并计算 1+2+...n的值，并存放到resChan
// 3) 最后8个协程协同完成工作后，再遍历resChan， 显示结果[如 res[1]=1... res[10]=55.)
// 4) 注意：考虑 resChan chan int 是否合适？
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func writeNum(numChan chan int) {
	for i := 1; i <= 2000; i++ {
		numChan <- i
		fmt.Printf("Write numChan: %v\n", i)
	}
	close(numChan)
}
func readCaluateNum(numChan chan int, resChan chan int) {
	defer wg.Done()
	for {
		v, ok := <-numChan
		if !ok {
			break
		}
		res := 0
		for i := 1; i <= v; i++ {
			res += i
		}
		fmt.Printf("Read: %v ,Caluate res: %v\n", v, res)
		resChan <- res

	}

}
func main() {

	numChan := make(chan int, 50)
	resChan := make(chan int, 2000)

	go writeNum(numChan)
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go readCaluateNum(numChan, resChan)
	}
	// time.Sleep(time.Second * 5)
	wg.Wait()
	close(resChan)
	sum := 0
	for v := range resChan {
		fmt.Printf("ResChan value: %v\n", v)
		sum += 1
	}
	fmt.Printf("sum: %v\n", sum)

}
