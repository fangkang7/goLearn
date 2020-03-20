package main

import "fmt"

/**
兼并另一个切片
*/
func main() {
	var slice1 = []int{1, 2, 3}
	slice2 := []int{4, 5, 6, 7, 8, 9}
	slice1 = append(slice1, slice2...)
	print(len(slice1))
	print(cap(slice1))
	fmt.Println(slice1)
}
