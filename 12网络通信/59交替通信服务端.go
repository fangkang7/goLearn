package main

import (
	"fmt"
	"net"
	"os"
)

/**
处理异常
*/
func JtHandleError(err error, when string) {
	if err != nil {
		fmt.Println(when, "错误原因：", err)
		os.Exit(1)
	}
}
func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	JtHandleError(err, "net.Listen")
	fmt.Println("listener...")
	defer listener.Close()
	for {
		// 接入一个客户端
		conn, err := listener.Accept()
		JtHandleError(err, "listener.Accept")
		go JtIoWithConn(conn)
	}
}

func JtIoWithConn(conn net.Conn) {
	defer conn.Close()
	// 创建缓存区
	buffer := make([]byte, 1024)
	clientAddr := conn.RemoteAddr()
	for {
		// 读入消息
		n, err := conn.Read(buffer)
		JtHandleError(err, "conn.Read")
		// 客户端消息
		clientMsg := string(buffer[:n])
		// 打印出客户端发过来的消息
		fmt.Println(clientAddr, clientMsg)
		// 服务端给客户端写的消息
		conn.Write([]byte("服务端以阅：" + clientMsg))
		if clientMsg == "over" {
			fmt.Printf("%v以下线\n", clientAddr)
			break
		}
	}

}
