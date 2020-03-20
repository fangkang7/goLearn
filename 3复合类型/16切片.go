package main

import "fmt"

/**
相当于长度可以扩张的数组
经过数组截取的片段就是切片
*/
func main01() {
	// 定义一个十位的数组
	var array = [10]int{0, 23, 14, 56, 89, 65, 32, 14, 10, 23}
	// 截取数组的前9位
	slice := array[0:9]
	fmt.Printf("array的数据类型是%T\n", array)
	fmt.Printf("slice的数据类型是%T,值是%v\n", slice, slice)

	// 下标为0开始截取5位
	slice = array[0:5]
	fmt.Printf("slice的数据类型是%T,值是%v\n", slice, slice)

	// 从0开始截取5位
	slice = array[:5]
	fmt.Printf("slice的数据类型是%T,值是%v\n", slice, slice)

	// 从下标0开始截取全部
	slice = array[0:]
	fmt.Printf("slice的数据类型是%T,值是%v\n", slice, slice)

	// 截取全部
	slice = array[:]
	fmt.Printf("slice的数据类型是%T,值是%v\n", slice, slice)

}

func main() {
	main01()
	var slice []int = []int{}
	fmt.Printf("slice的数据类型是%T\n", slice)
	fmt.Printf("slice的长度是%d,值是%v\n", len(slice), slice)

	// 往切片里边添加元素
	slice = append(slice, 0)
	fmt.Printf("slice的长度是%d,值是%v\n", len(slice), slice)

	slice = append(slice, 11, 22, 33)
	fmt.Printf("slice的长度是%d,值是%v\n", len(slice), slice)

	// 在来熟悉一下遍历数组
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice的第%d个元素的值是%v\n", i, slice[i])
	}

	for index, value := range slice {
		fmt.Printf("slice的第%d个元素的值是%v\n", index, value)
	}

}
