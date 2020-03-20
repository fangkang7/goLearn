package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

/**
os标准库
*/
func main010() {
	// 获得工作路径：当前工程根目录
	dir, err := os.Getwd()
	// C:\Users\Administrator\Desktop\go
	fmt.Println(dir)
	// <nil>
	fmt.Println(err)

	// 获取环境变量
	goroot := os.Getenv("GOROOT")
	// C:\Go
	fmt.Println(goroot)
}

/**
time标准库
*/
func main() {
	// 获取当前时间
	nowTime := time.Now()
	// 20-03-17 17:50:43.6997161 +0800 CST m=+0.005984401
	fmt.Println(nowTime)
	// 获取当前年份
	year := nowTime.Year()
	// 2020
	fmt.Println(year)
	// 获取当前月份
	month := nowTime.Month()
	// March
	fmt.Println(month)
	// 获取当前时间戳
	unixTime := time.Now().Unix()
	// 1584439175
	fmt.Println(unixTime)

	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("failure", "u & i"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))
}
