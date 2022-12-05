package main

import "fmt"

type Monkey struct {
	Name string
}

func (this *Monkey) climbing() {
	fmt.Println(this.Name, "会爬树")
}

// 声明接口，继承是补充
type BirdAble interface {
	Flying()
}
type LittleMonkey struct {
	Monkey
}

func (this *LittleMonkey) Flying() {
	fmt.Println(this.Name, "学习会飞")
}
func main() {
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}
	monkey.climbing()
	monkey.Flying()
}

// 当A结构体继承了B结构体，那么A结构体自动继承了B结构体的字段和方法
// 当A结构体需要扩展，那么通过接口补充
