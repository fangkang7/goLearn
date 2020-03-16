package main

import "fmt"

func main() {
	// 创建数组 空数组
	//var array [5]int
	// 创建数组的几种方式
	//var array = [5]int{3, 1, 4}
	// 长度可知的数组定义方式
	//array := [5]int{3, 1, 4, 5, 7}
	// 长度未知的数组定义方式
	array := [...]int{3, 1, 4, 7, 25, 8, 6, 9, 9}
	fmt.Println(array)
	// 获取数组长度
	fmt.Println(len(array))
	// 根据数组下标获取值
	fmt.Println("第一个元素是", array[0])

	// 修改数组的值
	array[0] = 123
	fmt.Println("第一个元素是", array[0])

	// 数组遍历
	for i := 0; i < len(array); i++ {
		fmt.Printf("数组的第%d个元素是%v\n", i, array[i])
	}

	for value := range array {
		fmt.Println(value)
	}

	for index, value := range array {
		fmt.Println(index, value)
	}
}
