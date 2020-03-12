package main

import "fmt"

/**
常量定义跟PHP一样使用const来定义
 */
// 可以不用声明类型
const pi = 3.14;
// 声明类型为float64
const pai float64  = 3.1415926;

/**
多个const值定义
 */
 const (
 	k = 123;
 	a = 456
 );

func main() {
	fmt.Println(k)
}