
// package main

// import "fmt"

// func add() func(int) int {
// 	var x int
// 	return func(i int) int {
// 		x += i
// 		return x
// 	}

// }
// func cal(base int) (func(int) int, func(int) int) {
// 	add := func(a int) int {
// 		base += a
// 		return base
// 	}
// 	sub := func(b int) int {
// 		base -= b
// 		return base
// 	}
// 	return add, sub
// }
// func main() {
// 	var f = add()
// 	fmt.Println(f(10))
// 	fmt.Println(f(30))
// 	f1 := add()
// 	r := f1(100)
// 	fmt.Printf("r: %v\n", r)
// 	c := f1(200)
// 	fmt.Printf("r: %v\n", c)
// 	add1, sub1 := cal(10)
// 	fmt.Printf("add1(1): %v\n", add1(1))
// 	fmt.Printf("sub1(2): %v\n", sub1(2))
// }