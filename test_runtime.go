// // 让出cpu时间片，重新等待安排任务
// package main

// import (
// 	"fmt"
// 	"runtime"
// )

//	func show(msg string) {
//		for i := 0; i < 2; i++ {
//			fmt.Printf("msg:%v\n", msg)
//		}
//	}
//
//	func main() {
//		go show("java")
//		for i := 0; i < 2; i++ {
//			runtime.Gosched() //我有权利执行任务了，让给你（其他子协程执行）
//			fmt.Printf("golang :%v\n","golang")
//		}
//	}
package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("a:", i)
	}
}
func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("b:", i)
	}
}
func show(msg string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v:msg:%v\n", i, msg)
		if i >= 5 {
			runtime.Goexit() //退出进程
		}
	}
}
func main() {
	// go show("java")
	// time.Sleep(time.Second)
	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	go a()
	go b()
	time.Sleep(time.Second * 3)
	fmt.Println("end ...")
}
