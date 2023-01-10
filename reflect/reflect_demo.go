package main

import (
	"fmt"
	"reflect"
)

func reflecrTest01(b interface{}) {
	// 通过反射获得传入变量的type,kind,值
	// 1 reflect.Type取得类型
	rTyp := reflect.TypeOf(b)
	fmt.Printf("rTyp: %v\n", rTyp)
	// 2 获取到reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal: %v rVal type: %T\n", rVal, rVal)
	//	n1 := 10
	//	n2 := rVal.(int) + n1
	//
	// rVal 不能直接计算
	iv := rVal.Interface()
	num2 := iv.(int)
	fmt.Printf("num2: %v %T\n", num2, num2)
}

// 对结构体反射
func reflecrTest02(c interface{}) {
	rVal := reflect.ValueOf(c)
	rType := reflect.TypeOf(c)
	iv := rVal.Interface()
	fmt.Printf("iv: %v iv type: %T\n", iv, iv)
	fmt.Printf("iv.(student1).Age: %v\n", iv.(student1).Age)
	// 获取Kind
	fmt.Printf("rVal.Kind(): %v\n", rVal.Kind())
	fmt.Printf("rType.Kind(): %v\n", rType.Kind())
}

type student1 struct {
	Name string
	Age  int
}

func main() {
	// var num int = 100
	// reflecrTest01(num)
	stu := student1{
		Name: "zhang",
		Age:  10,
	}

	reflecrTest02(stu)
	const tax int = 0
	fmt.Printf("tax: %v\n", tax)
	const (
		a = iota
		b
		c
		d
	)
	fmt.Println("a,b,c,d:", a, b, c, d)
	var num1 int = 19
	rVal := reflect.ValueOf(&num1)
	// rVal.SetInt(20)
	// fmt.Printf("num1: %v\n", num1)

	fmt.Printf("rVal.CanSet(): %v\n", rVal.Elem().CanSet())

}
