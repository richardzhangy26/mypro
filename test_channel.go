package main

import (
	"fmt"
	"math/rand"
	"time"
)

var values = make(chan int)

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Printf("value: %v\n", value)
	values <- value
	time.Sleep(time.Second * 5)
}
func main() {
	defer close(values)
	go send()
	fmt.Println("wait...")
	value := <-values
	fmt.Printf("receive:%v\n", value)
	fmt.Println("end..")
}
