package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

/**
文件打开
*/
func main0() {
	file, e := os.Open("D:/kaka.txt")
	if e != nil {
		fmt.Println("文件打开失败")
	} else {
		fmt.Println("文件打开成功")
	}

	defer func() {
		file.Close()
		fmt.Println("文件已关闭")
	}()
	fmt.Println("愉快滴欣赏avi艺术")
	time.Sleep(3 * time.Second)
}

/**
文件打开模式
*/
func main1() {
	/*
		+ O_RDONLY int = syscall.O_RDONLY // 只读模式
		+ O_WRONLY int = syscall.O_WRONLY // 只写模式
		+ O_RDWR   int = syscall.O_RDWR   // 读写模式
		+ O_APPEND int = syscall.O_	APPEND // 追加模式
		+ O_CREATE int = syscall.O_CREAT  // 文件不存在则创建
		+ O_TRUNC  int = syscall.O_TRUNC  // 覆盖模式
	*/
}

/**
文件读取
*/
func main2() {
	file, e := os.OpenFile("D:/kaka.txt", os.O_RDONLY, 0666)
	if e != nil {
		fmt.Println("文件打开失败：err", e)
		return
	} else {
		fmt.Println("文件打开成功")
	}

	// 打开时立刻挂起关闭程序
	defer func() {
		file.Close()
		fmt.Println("文件关闭成功")
	}()

	// 创建文件的读取器
	reader := bufio.NewReader(file)
	// 死循环读取
	for {
		// 以换行符为界，分批次读取数据，得到readString
		//readString, e := reader.ReadString('\n')
		line, _, e := reader.ReadLine()
		if e != nil {
			fmt.Println("文件读取失败;err", e)
			// 如果已到文件末尾，则直接结束
			if e == io.EOF {
				time.Sleep(time.Second)
				break
			}
		}
		//fmt.Print(readString)
		fmt.Println(string(line))
	}
	fmt.Println("读取结束")
}

/**
简易的数据读取
*/
func main3() {
	bytes, err := ioutil.ReadFile("d:/kaka.txt")
	if err != nil {
		fmt.Println("读取失败，err=", err)
		return
	} else {
		contentStr := string(bytes)
		fmt.Println(contentStr)
	}
}

/**
关于文件读取少了一行的另一种解决方案
*/
func main() {
	file, e := os.Open("d:/kaka.txt")
	if e != nil {
		fmt.Println("文件打开失败")
	} else {
		fmt.Println("文件打开成功")
		// 创建文件读取器
		reader := bufio.NewScanner(file)
		for reader.Scan() {
			fmt.Println(reader.Text())
		}
	}
}
