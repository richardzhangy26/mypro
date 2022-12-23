package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 390
	close(intChan)
	// close后只能取数据，不能存数据
	intChan2 := make(chan int, 100)
	go func() {
		for i := 0; i < 100; i++ {
			intChan2 <- i * 2
		}
		close(intChan2)
	}()
	//普通遍历会使得长度减少，最终遍历只有一半
	// for i := 0; i < len(intChan2); i++ {
	// 	num := <-intChan2
	// 	fmt.Printf("i: %v num: %v cap(channel): %v, len(channel) %v\n", i, num, cap(intChan2), len(intChan2))
	// }
	// fmt.Printf("intChan2: %T\n", intChan2)
	// close(intChan2)
	// for v := range intChan2 {
	// 	fmt.Printf("v: %v\n", v)

	// }

	for {
		v, ok := <-intChan2
		if !ok {
			fmt.Printf("ok: %v v: %v\n", ok, v)
			break
		}
		fmt.Printf("ok: %v v: %v\n", ok, v)
	}

}
