package main

import (
	"fmt"
	"net"
	"os"
)

func handleError(err error, when string) {
	if err != nil {
		fmt.Println("错误地址,err", when, err)
		// 非正常退出,正常退出为0
		os.Exit(0)
	}
}

func main() {
	// 拨号请求连接
	conn, e := net.Dial("udp", "127.0.0.1:8888")
	handleError(e, "Dial")
	defer conn.Close()

	// 向连接中写出信息
	n, err := conn.Write([]byte("咔咔手赚网地址：fangkang.top"))
	handleError(err, "Write")
	fmt.Println("写出字节数", n)

	// 创建缓冲区
	buffer := make([]byte, 1024)
	n, err = conn.Read(buffer)
	handleError(err, "conn.Read")
	fmt.Println(string(buffer[:n]))
}
