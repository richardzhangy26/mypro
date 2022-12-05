package main

import "github.com/gin-gonic/gin"

func hello(c *gin.Context) {
	c.String(200, "hello %s", "zyc")
}

func main() {
	e := gin.Default()
	e.GET("/hello", hello)
	e.Run(":8989")
}

//create a blockchian
