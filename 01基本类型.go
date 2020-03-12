package main

import "fmt"

func main()  {
	fmt.Printf("数据类型是:%T\n", 123)
	fmt.Printf("数据类型是:%T\n", 123.4)
	fmt.Printf("数据类型是:%T\n", "咔咔博客blog.fangkang.top")
	fmt.Printf("数据类型是:%T\n", false)
	// 指针
	var a int = 123;
	fmt.Printf("数据类型是:%T\n", &a)
	fmt.Printf("数据类型是:%V\n", &a)
	// 固定5长度的整形数组
	fmt.Printf("数据类型是:%T\n", [5]int{1,2,3,4,5})
	// 可以动态扩容的整形切片
	fmt.Printf("数据类型是:%T\n", []int{1,2,3,4,5})
	// 映射
	var scores map[string]int = make(map[string]int)
	scores["语文"] = 123;
	scores["数学"] = 456;
	fmt.Printf("数据类型是:%T\n", scores)
	fmt.Println(scores["语文"])
	// 结构体
	type person struct {
		name string
		age  int
		rmb float64
	}
	var p person = person{name:"咔咔",age:20,rmb:12.34}
	fmt.Printf("数据类型是%T",p)
	var a int
}
