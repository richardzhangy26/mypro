package main

import "fmt"

func main() {
	var a *int
	a = new(int)
	*a = 5
	fmt.Printf("a:%T\n %v\n", a, &*a)
	var b *[]int
	b = new([]int)
	fmt.Printf("b: %T\n%v\n", b, b)
	var c []int = make([]int, 2)
	fmt.Printf("c: %T\n%v\n", c, &c)
	c = append(c, 10)
	fmt.Printf("c: %T\n%v\n", c, c)
	fmt.Printf("c: %v\n", &c)
	var d *[]int
	d = &c
	fmt.Printf("d: %v\n", d)
}
