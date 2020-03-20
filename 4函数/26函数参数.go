package main

import "fmt"

func main() {
	//sayHello("kaka", 10)
	//Sum(12, 35)
	//Sum1(12, 35)
	Sum2(3, 4, 5, 6)
}

/**
不定长的参数
*/
func Sum2(a ...int) {
	// a的类型是[]int,a的值是[3 4 5 6]
	fmt.Printf("a的类型是%T,a的值是%v\n", a, a)
	result := 0
	//for i := 0; i < len(a); i++ {
	//	result += a[i]
	//}
	// 这里需要介绍一个点那就是当我们使用不到index时，可以使用_来代替
	for _, value := range a {
		result += value
	}
	// 所有值的和为 18
	fmt.Println("所有值的和为", result)
}

/**
同种类型参数合并声明
*/
func Sum(a int, b int) {
	fmt.Printf("a与b的和是%d\n", a+b)
}

/**
同种类型参数合并声明
*/
func Sum1(a, b int) {
	fmt.Printf("a与b的和是%d\n", a+b)
}

/**
多个参数
切记参数需要传入类型  这里跟PHP有点出入
*/
func sayHello(name string, count int) {
	for i := 0; i < count; i++ {
		fmt.Println("hello", name)
	}
}
