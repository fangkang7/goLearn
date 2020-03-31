package main

import (
	"fmt"
	"net"
	"os"
)

func TcpError(err error, when string) {
	if err != nil {
		fmt.Println("错误地址,err", when, err)
		// 非正常退出,正常退出为0
		os.Exit(0)
	}
}

func ioWithConn(conn net.Conn) {
	// 开辟一个缓冲区
	buffer := make([]byte, 1024)
	// 关闭连接
	defer conn.Close()
	for {
		// 读取数据
		n, err := conn.Read(buffer)
		// 验证
		TcpError(err, "conn.Read")
		// 把读取的数据放到缓存区中
		clientMsg := string(buffer[:n])
		fmt.Println("收到客户端消息", clientMsg)
		// 如果客户端输入的是 im off
		if clientMsg == "im off" {
			// 服务端往连接写一个bye
			conn.Write([]byte("bye"))
			break
		}
		// 服务端回复客户端的信息
		conn.Write([]byte("msg receive" + clientMsg))
	}
	fmt.Println("用户断开连接")
}

func main() {
	// 监听tcp连接
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	// 验证错误
	TcpError(err, "net.Listen")
	// 服务端一连接就会出现listener...
	fmt.Println("listener...")
	for {
		// 新的客户端连接
		conn, err := listener.Accept()
		// 验证连接
		TcpError(err, "listener.Accept")
		// 处理每一个客户端
		go ioWithConn(conn)
	}

}
