package main

import "fmt"

// var i int = initvar()

func init() {
	fmt.Println("init start")
}
func initvar() int {
	fmt.Println("initvar start")
	return 100
}
func main() {
	fmt.Println("main start")
}

//先执行变量init再执行init函数最后是main
