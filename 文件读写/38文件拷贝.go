package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/**
使用io的copy接口
*/
func main6() {
	// 原文件 读取文件需要读的权限
	srcfile, _ := os.OpenFile("d:/kaka.txt", os.O_RDONLY, 0666)
	// 目标文件  需要写和创建权限
	disfile, _ := os.OpenFile("d:/kaka123.txt", os.O_WRONLY|os.O_CREATE, 0666)

	// 挂起文件，结束时关闭文件
	defer func() {
		srcfile.Close()
		disfile.Close()
	}()
	// 拷贝成功，返回成功的字节数
	written, err := io.Copy(disfile, srcfile)
	if err == nil {
		// 文件拷贝成功，字节数： 72
		fmt.Println("文件拷贝成功，字节数：", written)
	} else {
		fmt.Println("文件拷贝失败")
	}
}

/**
最简单的copy案例
*/
func main5() {
	bytes, _ := ioutil.ReadFile("d:/kaka.txt")
	err := ioutil.WriteFile("d:/kakas.txt", []byte(bytes), 0666)
	if err != nil {
		fmt.Println("拷贝失败")
	} else {
		fmt.Println("拷贝成功")
	}
}

func main() {
	// 打开原始文件
	srcFile, _ := os.OpenFile("d:/kaka.txt", os.O_RDONLY, 0666)

	// 打开目标文件
	disFile, _ := os.OpenFile("d:/kakas.txt", os.O_CREATE|os.O_WRONLY, 0666)
	// 创建一个4字节的缓冲区
	buffer := make([]byte, 4)
	reader := bufio.NewReader(srcFile)
	writer := bufio.NewWriter(disFile)
	for {
		// 使用缓存区来读取数据
		n, err := reader.Read(buffer)
		fmt.Println(n, err)
		// 将缓冲区的数据读出到目标文件
		writer.Write(buffer)
		// 读到文件末尾，清空最后一组数据，退出循环
		if err == io.EOF {
			writer.Flush()
			break
		}
	}
	fmt.Println("拷贝完成")
}
