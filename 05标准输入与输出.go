package main

import "fmt"

func main0() {
	// 定义一个变量  复制为  blog.fangkang.top  系统就会开辟一片内存空间 获取kakaBlog的内存地址
	var kakaBlog string = "blog.fangkang.top"
	// 阻塞 接收用户的输入的值 存入kakaBlog的内存地址
	// &取出kakaBlog的地址 &kakaBlog取出kakaBlog所在的地址
	fmt.Scan(&kakaBlog)
	// 打印kakaBlog的值
	fmt.Println(kakaBlog)
}

const pi = 3.14

func main1() {
	var r = 0.0
	fmt.Println("咔咔请输出需要计算的半径：")
	fmt.Scan(&r)
	// 同类型才可以做计算   pi是float类型  如果r定义为int类型是不可以做计算的
	area := pi * r * r
	fmt.Printf("面积为%v\n", area)
}

func main() {
	// 声明变量可以分开设置，也可以一起设置  一起设置需要类型一致
	//var nidegongzi int
	//var wifegongzi int
	var nigongzi, wifegognzi int
	fmt.Println("请输入你和你妻子的工资")
	// 一次性接收多个输入，使用空格或者回车隔开  分别丢入相应的内存地址中
	fmt.Scan(&nigongzi, &wifegognzi)
	gognzi := (nigongzi + wifegognzi) * 12
	fmt.Printf("你们家一年的工资收入是%v\n", gognzi)
}
