package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	var allChan chan interface{}
	allChan = make(chan interface{}, 5)
	// catChan := make(chan interface{})
	cat1 := Cat{Name: "shi", Age: 10}
	cat2 := Cat{Name: "ershi", Age: 20}
	allChan <- cat1
	allChan <- cat2
	allChan <- 10
	allChan <- "jack"
	cat11 := <-allChan
	fmt.Printf("cat11: %v\n", cat11.(Cat).Name)
	fmt.Printf("cat11: %T\n", cat11)
	<-allChan
	num := <-allChan
	fmt.Printf("(num + 5): %v\n", (num.(int) + 5))
	fmt.Printf("num: %T\n", num.(int))

}
