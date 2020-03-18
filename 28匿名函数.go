package main

import (
	"fmt"
	"time"
)

/**
认识匿名函数
*/
func main01() {
	// 匿名函数
	sum := func(a, b int) int {
		return a + b
	}(3, 4)
	fmt.Println(sum)
}

/**
defer的应用场景
*/
func main02() {
	// 这里的defer只是一个简单的输出，假如是很多的业务逻辑需要处理
	// 但又不需要单独封装的话，就需要使用defer来做这件事情喽！
	//defer fmt.Println("我是最后一个执行的")

	defer func(count int) {
		for i := 0; i <= count; i++ {
			fmt.Println("咔咔手赚网的地址是：fangkang.top")
			fmt.Println("咔咔博客的地址是：blog.fangkang.top")
		}
	}(3)
	// 在之前的学习过程中呢！我们学习过这个defer指的是延迟执行
	fmt.Println("一二三四五")
	fmt.Println("咔咔手赚网")
	fmt.Println("咔咔博客")
	fmt.Println("fangkang.top")
}

/**
go的应用场景
*/
func main() {
	// go就是并发执行
	go func() {
		for i := 2; i <= 10; i += 2 {
			fmt.Println("我是并发协程", i)
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println(i)
		// 等待秒钟执行
		time.Sleep(time.Second)
	}
}
