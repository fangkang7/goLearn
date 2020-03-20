package main

import "fmt"

func main() {
	var count int
	for i := 1; i < 10; i++ {
		count++
		fmt.Printf("咔咔第%v循环\n", i)
		if count == 3 {
			fmt.Println("咔咔的第三次不去执行")
		}
		if count >= 5 {
			fmt.Println("咔咔已经执行了第五次了")
			break
		}
	}
}
