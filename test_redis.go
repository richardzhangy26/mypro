package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println("Dial err = ", err)
		return
	}
	_, err = c.Do("set", "name", "tomjerry")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}
	r, err := redis.String(c.Do("get", "name"))
	if err != nil {
		fmt.Println("String err =", err)
		return
	}
	fmt.Printf("r: %v\n", r)

	// fmt.Println("conn redis sucessful")
	defer c.Close()

}
