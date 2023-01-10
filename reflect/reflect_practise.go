package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string  `json:"name"`
	Age   int     `json:"monster_age"`
	Score float32 `json:"成绩"`
	Sex   string  `json:"性别"`
}

func (s Monster) Print() {
	fmt.Println("=====start====")
	fmt.Println(s)
	fmt.Println("=====end=====")
}
func (s Monster) Getsum(n1 int, n2 int) int {
	return n1 + n2

}
func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}
func TestStruct(a interface{}) {
	rTyp := reflect.TypeOf(a)
	rVal := reflect.ValueOf(a)
	kd := rVal.Kind()
	if kd != reflect.Struct {
		fmt.Println("不是结构体返回")
		return
	}
	num := rVal.NumField()
	fmt.Printf("rval的字段数量: %v\n", num)
	for i := 0; i < num; i++ {
		fmt.Printf("rVal.Field(%v): %v\n", i, rVal.Field(i))
		fmt.Printf("rTyp.Field(%v).Tag.Get(\"json\"): %v\n", i, rTyp.Field(i).Tag.Get("json"))
	}
	fmt.Printf("rVal.NumMethod(): %v\n", rVal.NumMethod())
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(30))
	res := rVal.Method(0).Call(params)
	fmt.Printf("res[0].Int(): %v\n", res[0].Int())

}

func main() {
	var a Monster = Monster{
		Name:  "黄鼠狼",
		Age:   400,
		Score: 99.5,
	}
	TestStruct(a)
}
