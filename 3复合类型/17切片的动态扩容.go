package main

import "fmt"

/**
cap(slice)获取切片的容器
切片创建初期，容器等于长度
动态扩张时，一旦容量无法满足，就会以翻倍的策略进行扩张
*/
func main() {
	var slice = []int{1}
	fmt.Printf("slice的长度是%d,容量是%d\n", len(slice), cap(slice))
	slice = append(slice, 4)
	fmt.Printf("slice的长度是%d,容量是%d\n", len(slice), cap(slice))
	slice = append(slice, 5, 6)
	fmt.Printf("slice的长度是%d,容量是%d\n", len(slice), cap(slice))
	slice = append(slice, 5, 6)
	fmt.Printf("slice的长度是%d,容量是%d\n", len(slice), cap(slice))
}
