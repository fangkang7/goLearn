package main

import "fmt"

var hero = []string{"关羽", "张飞", "赵云", "马超", "黄忠"}

func main() {
	nextHero1 := NextHere1()
	for i := 0; i < 7; i++ {
		// 关羽 张飞 赵云 马超 黄忠 关羽 张飞
		fmt.Println(nextHero1())
	}
}

/**
闭包函数：返回函数的函数
*/
func NextHere1() func() string {
	var index int
	f := func() string {
		// 获取到对应的值
		hero := hero[index]
		// 每次加1
		index++
		// 判断当索引大于4的时候把索引的初始值在改为0
		if index > 4 {
			index = 0
		}
		return hero
	}
	return f
}

/**
练习闭包写法
这个错误是索引超出了范围
panic: runtime error: index out of range [5] with length 5
*/
func NextHere2() func() string {
	var index int
	res := func() string {
		s := hero[index]
		index++
		if index > 4 {
			index = 0
		}
		return s
	}
	return res
}

/**
这个错误是索引超出了范围
panic: runtime error: index out of range [5] with length 5
*/
//func NextHero() string {
//	// 获取到对应的值
//	hero := hero[index]
//	// 每次加1
//	index++
//	// 判断当索引大于4的时候把索引的初始值在改为0
//	if index > 4 {
//		index = 0
//	}
//	return hero
//}
