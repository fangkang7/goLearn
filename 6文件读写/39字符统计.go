package main

import "fmt"

/**
字符串的遍历
*/
func main7() {
	str := "咔咔博客地址：blog.fangkang.top"
	for i, v := range str {
		fmt.Printf("序号%d,%c,类型%T\n", i, v, v)
	}
}

/**
字节和字符串互转
*/
func main8() {
	// 字符串转字节
	// [229 146 148 229 146 148]
	fmt.Println([]byte("咔咔"))

	// 字节转字符串
	// 咔咔
	fmt.Println(string([]byte{229, 146, 148, 229, 146, 148}))
}

func main() {
	main7()
	main8()
}
