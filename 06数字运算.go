package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(2 + 3)
	fmt.Println(2 - 3)
	fmt.Println(2 * 3)
	fmt.Println(2 / 3)
	// 求乘方
	result := math.Pow(2, 3)
	fmt.Printf("%v的%v的乘方是%v\n", 2, 3, result)

	// 求平方根
	result = math.Sqrt(9)
	fmt.Printf("%v的根是%v\n", 9, result)

	// 求绝对值
	result = math.Abs(-10)
	fmt.Println(result)
	fmt.Printf("%v的绝对值是%v\n", -10, result)
	// 纯舍(舍掉小数)
	result = math.Floor(3.6)
	fmt.Println(result)
	fmt.Printf("%v的纯舍是%v\n", 3.6, result)
	// 纯入(直接给原数加1)
	result = math.Ceil(3.01)
	fmt.Println(result)
	fmt.Printf("%v的纯入是%v\n", 3.01, result)
	// 四舍五入
	result = math.Round(3.4)
	fmt.Println(result)
	fmt.Printf("%v的四舍五入是%v\n", 3.4, result)
}
