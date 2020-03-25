package main

import "fmt"

func main() {
	fmt.Println(GetFibonacci(6))
	sum := GetNum(10)
	fmt.Println(sum)
	getRecursion := GetRecursion(10)
	fmt.Println(getRecursion)
	fbNum := GetFbNum(10)
	fmt.Println(fbNum)
}

/**
递归就是自己调自己
递归一定要有终止条件，否则就会无限循环
*/
func GetFibonacci(n int) int {
	// 如果是第0项或者第2项直接返回1
	if n == 0 || n == 1 {
		return 1
	} else {
		return GetFibonacci(n-1) + GetFibonacci(n-2)
	}
}

/**
使用非递归实现斐波那契数列
*/
func GetFbNum(n int) int {
	a := 1
	b := 1
	c := a + b
	for i := 1; i <= n; i++ {
		a = b
		b = c
		c = a + b
	}
	return a
}

/**
使用循环来实现自然数之和
*/
func GetNum(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += i
	}
	return
}

/**
使用递归来实现自然数求和
*/
func GetRecursion(n int) (sum int) {
	if n == 1 {
		return 1
	} else {
		return n + GetRecursion(n-1)
	}
}
