package main

import "fmt"

type BInterface interface {
	test02()
}
type CInterface interface {
	test03()
}

// 接口继承
type AInterface interface {
	test01()
	BInterface
	CInterface
}
type Stu struct {
	name string
}

// 空接口
type T interface {
}

func (stu *Stu) test01() {
	fmt.Println("this is test01")

}
func (stu *Stu) test02() {
	fmt.Println("this is test02")
}
func (stu *Stu) test03() {
	fmt.Println("this is test03")
}
func main() {
	var stu Stu
	stu.name = "zhangyichi"
	var A AInterface = &stu
	A.test01()
	var t T = stu
	fmt.Printf("t: %v\n", t)
	// 空接口能接受任意类型
	var t2 interface{}
	var num float64 = 8.8
	t2 = num
	fmt.Printf("t2: %v\n", t2)

}
