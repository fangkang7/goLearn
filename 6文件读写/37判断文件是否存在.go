package main

import (
	"fmt"
	"os"
)

func main() {
	checkFile("D:/kaka.txt")
}

func checkFile(filepath string) (exist bool) {
	// 是一个文件的指针
	fileInfo, e := os.Stat(filepath)
	// &{kaka.txt 32 {1848971300 30801167} {1194119857 30801333} {1194119857 30801333} 0 216 0 0 {0 0} d:/kaka.txt 0 0 0 false}
	//fmt.Println(fileInfo)
	if fileInfo != nil && e == nil {
		exist = true
		fmt.Println("文件存在")
		// 判断文件是否不存在
	} else if os.IsNotExist(e) {
		fmt.Println("文件不存在")
		exist = false
	}
	// 这里返回可以带返回值的名，也可以不带
	return
	//return exist
}
