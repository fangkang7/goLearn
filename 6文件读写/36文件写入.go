package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//覆写模式
	file, e := os.OpenFile("c:/kaka.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0754)
	//追加模式
	//file, e := os.OpenFile("d:/kaka.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0754)
	// 创写模式
	//file, e := os.OpenFile("c:/kaka.txt", os.O_CREATE|os.O_WRONLY, 0777)
	if e != nil {
		fmt.Println("文件打开失败")
		return
	} else {
		fmt.Println("文件打开成功")
	}
	// 挂起关闭程序
	defer file.Close()
	// 创建写入器
	writer := bufio.NewWriter(file)
	// 执行带缓冲的写入
	writer.WriteString("咔咔博客地址\n")
	writer.WriteString("blog.fangkang.top\n")
	writer.WriteString("咔咔手赚网地址\n")
	writer.WriteString("fangkang.top\n")
	// 强制将缓冲区中的内容写入文件
	writer.Flush()
	fmt.Println("文件写入完毕")
}

func main4() {
	file := ioutil.WriteFile("d:/kaka.txt", []byte("咔咔你好"), 0666)
	if file != nil {
		fmt.Println("文件写入成功")
	} else {
		fmt.Println("文件写入成功")
	}

}
