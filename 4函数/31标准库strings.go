package main

import (
	"fmt"
	"strings"
)

func main() {
	// 判断一个字符串是否在另一个字符串中
	contains := strings.Contains("substr", "s")
	// true
	fmt.Println(contains)

	// 判断一个字符串是否存在另一个字符串的任意字符
	any := strings.ContainsAny("kaka", "nihao")
	// true
	fmt.Println(any)

	// 判断一个字符串在另一个字符串第一次出现的位置  注意空格也算一个字符
	index := strings.Index("Hello World", "or")
	// 7
	fmt.Println(index)

	// 把一个字符串转小写
	lower := strings.ToLower("KAKA")
	// kaka
	fmt.Println(lower)

	// 把一个字符串转大写
	upper := strings.ToUpper("kaka")
	// KAKA
	fmt.Println(upper)

	// 把字符串分割为切片，根据某一标识
	after := strings.SplitAfter("a,b,c", ",")
	// 类型为[]string,值为[a, b, c]
	fmt.Printf("类型为%T,值为%v\n", after, after)
}
