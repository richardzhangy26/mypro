// package main

// import (
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	ginServer := gin.Default()
// 	ginServer.LoadHTMLFiles("templates/index.html")
// 	ginServer.Static("/static", "./static")
// 	ginServer.GET("/hello", func(ctx *gin.Context) {
// 		ctx.HTML(200, "index.html", gin.H{"msg": "hello richradzhang"})
// 	})

// 	// ginServer.POST("/user",func(ctx *gin.Context) {
// 	// 	ctx.JSON()
// 	// })

//		ginServer.Run("127.0.0.1:8082")
//	}
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
	fmt.Print(a, z)
}
