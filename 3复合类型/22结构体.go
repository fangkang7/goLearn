package main

import "fmt"

func main() {
	type Person struct {
		name string
		age  int
		sex  bool
		like []string
	}
	p := Person{}
	p.name = "咔咔"
	p.age = 24
	p.sex = true
	p.like = []string{"buk", "写代码"}

	// 直接获取值
	fmt.Printf("p的值是%v\n", p)
	// 获取键值
	fmt.Printf("p的值是%+v\n", p)
	// 获取结构体和键值
	fmt.Printf("p的值是%#v\n", p)
}
