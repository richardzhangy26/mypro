package main

import "fmt"

type Animal interface {
	// Speak() string
	eat()
}

type Dog struct {
	name string
}

func (dog Dog) eat() {
	fmt.Printf("dog: %v\n", dog)
	fmt.Println("dog eat ----")
	dog.name = "嘿嘿"
}

// func (dog *Dog) eat() {

func main() {
	dog := Dog{
		name: "花花",
	}
	dog1 := Animal
	dog.eat()
	fmt.Printf("dog: %v\n", dog)
}
