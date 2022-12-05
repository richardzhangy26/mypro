package main

import "fmt"

func printArray(args []interface{}) {
	for _, v := range args {
		fmt.Println(v)
	}
}
func printa(arg interface{}) {
	fmt.Println(arg)
}

func main() {
	// any := make([]interface{}, 5)
	// any[0] = 11
	// any[1] = "hello world"
	// any[2] = []int{11, 22, 33, 44}
	// for _, value := range any {
	// 	fmt.Println(value)
	// }
	// printArray(any)
	var arg interface{}
	arg = 1
	printa(arg)

}
