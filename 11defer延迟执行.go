package main

import (
	"fmt"
	"time"
)

func main() {

	defer fmt.Println("博客地址：blog.fangkang.top")
	defer fmt.Println("博客手赚网fangkang.top")
	defer fmt.Println("百度搜索咔咔手赚网")

	fmt.Println("咔咔博客")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("咔咔手赚网")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("手赚网地址：fangkang.top")
}
