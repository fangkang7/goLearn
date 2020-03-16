package main

import "fmt"

func main() {
	// 使用make直接定义
	slice := make([]int, 0)
	fmt.Println(slice, len(slice), cap(slice))

	// 直接var一个切片
	var slice1 []int = []int{}
	fmt.Println(slice1)
	fmt.Printf("slice1的类型是%T\n", slice1)

	// 通过对数组进行截取来获取切片
	var array [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("array的类型是%T\n", array)
	slice2 := array[:]
	fmt.Printf("array的类型是%T\n", slice2)
}
