package main

import "fmt"

type Person struct {
	name string
}

type Customer struct {
	name string
}

func (cus Customer) login(name string, pwd string) bool {
	fmt.Printf("cus.name: %v\n", cus.name)
	if cus.name == "tom" && pwd == "123" {
		return true
	} else {
		return false
	}

}

func (per Person) eat() {
	fmt.Printf("%v,eat...\n", per.name)
}
func (per Person) sleep() {
	fmt.Printf("%v,sleep...", per.name)
}

func main() {
	per := Person{
		name: "Tom",
	}
	per.eat()
	per.sleep()
	cus := Customer{
		name: "tom",
	}
	b := cus.login("tom", "123")
	fmt.Printf("b: %v\n", b)
}
