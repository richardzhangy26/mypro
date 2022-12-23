package main

import (
	"fmt"
)

func printCat(done chan bool) {
	for i := 0; i < 100; i++ {
		fmt.Println("cat")
	}
	done <- true
}

func printDog(done chan bool) {
	<-done
	for i := 0; i < 100; i++ {
		fmt.Println("dog")
	}
	done <- true
}

func printFish(done chan bool) {
	<-done
	for i := 0; i < 100; i++ {
		fmt.Println("fish")
	}
}

func main() {
	done := make(chan bool)
	go printCat(done)
	go printDog(done)
	go printFish(done)
	<-done
}
