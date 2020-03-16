package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 以当前时间纳秒为种子，获得一个1000以内的数据
	myRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := myRand.Intn(1001)

	for {
		var cai int
		fmt.Println("咔咔请输入你的猜想")
		fmt.Scan(&cai)
		if cai > result {
			fmt.Println("猜的大了")
		} else if cai < result {
			fmt.Println("猜的小了")
		} else {
			fmt.Println("猜对了")
		}
	}
}
