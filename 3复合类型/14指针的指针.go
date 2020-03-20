package main

import "fmt"

func main() {
	// 定义x变量
	var x = 123
	// 定义int类型的指针
	var mPrt *int = &x
	// 声明并赋值
	mmPrt := &mPrt
	fmt.Println(mmPrt)
	fmt.Printf("mPrt的类型为%T\n", mmPrt)

	// 需求不打印x获取x的值
	fmt.Println(*mPrt)
	fmt.Println(**mmPrt)
	fmt.Println(*(*mmPrt))
}
