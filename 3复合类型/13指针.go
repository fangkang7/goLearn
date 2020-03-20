package main

import "fmt"

func main() {
	// 声明的实际变量
	var a int = 20
	// 指针变量
	var ip *int
	// 指针变量的存储地址
	ip = &a
	*ip = 13
	fmt.Printf("a变量的地址是%v\n", &a)
	fmt.Printf("a变量的值是%v\n", a)
	fmt.Printf("ip变量的地址是%v\n", ip)
	fmt.Printf("*ip变量的地址是%v\n", *ip)
}
