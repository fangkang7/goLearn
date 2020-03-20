package main

import "fmt"

func main() {
	//sum := GetSum(3, 4, 5, 6)
	//sum := GetSum1(3, 4, 5, 6)
	sum, isEven := GetSum2(3, 4, 5, 6)
	// sum的值为 18
	fmt.Println("sum的值为", sum)
	// sum的奇偶为 true
	fmt.Println("sum的奇偶为", isEven)
}

func GetSum2(a ...int) (sum int, isEven bool) {
	for _, v := range a {
		sum += v
	}
	if sum%2 == 0 {
		isEven = true
	}
	return
}

func GetSum1(a ...int) (sum int) {
	for _, v := range a {
		sum += v
	}
	return
}

func GetSum(a ...int) int {
	var sum int
	for _, v := range a {
		sum += v
	}
	// 18
	//fmt.Println(sum)
	return sum
}
