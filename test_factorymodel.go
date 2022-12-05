/*
*
工厂模式：为了解决golang结构体没有构造函数
比如在package model包里定义了Student的结构体，可以将其引入到其他包里但首字母小写时就不行。
工厂模式解决
*
*/
package main

import (
	"fmt"
	"mypro/user/model"
)

func main() {
	// student := model.Student{
	// 	Name:  "yichizhang",
	// 	Score: 100.0,
	// }
	// fmt.Printf("student: %v\n", student)
	var stu = model.NewStudent("tom", 99.5)
	fmt.Printf("stu: %v\n", stu)
	fmt.Printf("stu: %v\n%v\n% v\n", *stu, stu.Name, stu.GetScore())
}
