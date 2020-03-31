package main

import (
	"fmt"
	"net"
	"os"
)

/**
处理异常
*/
func JcHandleError(err error, when string) {
	if err != nil {
		fmt.Println(when, "错误原因：", err)
		os.Exit(1)
	}
}
func main() {
	// 拨号连接,得到专线conn
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	JcHandleError(err, "net.Dial")
	defer conn.Close()
	// 在专线中处理io
	ioInteractive(conn)
}

func ioInteractive(conn net.Conn) {
	defer conn.Close()
	// 在一条独立的协程里接收服务端的消息
	go func() {
		// 创建一个缓冲区
		buffer := make([]byte, 1024)
		// 循环读取
		for {
			// 从缓冲区读取消息
			n, err := conn.Read(buffer)
			JcHandleError(err, "conn.Read")
			serverMsg := string(buffer[:n])
			fmt.Println("服务端消息:", serverMsg)
		}
	}()

	// 在主协程里向服务端写消息
	var userInput string
	for {
		// 标准输入与输出
		fmt.Scan(&userInput)
		// 向服务端写出
		conn.Write([]byte(userInput))
		if userInput == "exit" {
			return
		}
	}
}
