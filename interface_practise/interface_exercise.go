package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 1.声明HEro结构体
type Hero struct {
	Name string
	Age  int
}

// 2.声明一个Hero结构体切片类型
type HeroSlice []Hero

func (hs HeroSlice) Len() int {
	return len(hs)
}

// Less方法就是决定你使用什么标准进行排序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age > hs[j].Age
}
func (hs HeroSlice) Swap(i, j int) {
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp

}
func main() {
	var intSlice = sort.IntSlice{0, 1, -1, 10, 90}
	// 1.冒泡排序
	intSlice.Len()
	sort.Ints(intSlice)
	fmt.Println(intSlice)
	var heros HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄～%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heros = append(heros, hero)
	}
	fmt.Println("~~~~~根据年龄排序前~~~~~~~")
	for _, v := range heros {
		fmt.Printf("v: %v\n", v)
	}
	sort.Sort(heros)
	fmt.Println("~~~~~根据年龄排序后~~~~~~~")
	for _, v := range heros {
		fmt.Printf("v: %v\n", v)
	}
}
