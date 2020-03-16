package main

import "fmt"

// 切片扩容时  所有元素的地址都重新迁移到另一片地址
func main() {
	// 定义一个长度为5的数组
	var array = [5]int{0, 1, 2, 3, 4}
	// 定义切片slice1
	slice1 := array[:]
	// 定义切片slice2
	slice2 := slice1[:]
	fmt.Println(array, slice1, slice2)
	// 取出地址
	fmt.Printf("%p,%p,%p\n", &array, &slice1, &slice2)
	// 获取下标为0元素的地址
	fmt.Printf("%p,%p,%p\n", &array[0], &slice1[0], &slice2[0])
	// 获取下标为1元素的地址
	fmt.Printf("%p,%p,%p\n", &array[1], &slice1[1], &slice2[1])
	// 当我们修改其中一个切片的值时  相同地址的元素都会改变
	array[0] = 123
	slice1[1] = 456
	slice2[2] = 789
	fmt.Println(array, slice1, slice2)
	// 给slice2的切片进行扩容会改变slice2的地址  那也就跟array和slice1没有了任何关系
	slice2 = append(slice2, 5)
	array[0] = 1
	slice1[1] = 2
	slice2[0] = 456
	fmt.Println(array, slice1, slice2)
}
