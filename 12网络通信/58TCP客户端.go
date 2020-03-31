package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func tcpError(err error, when string) {
	if err != nil {
		fmt.Println("错误地址,err", when, err)
		// 非正常退出,正常退出为0
		os.Exit(0)
	}
}
func main() {
	// 直接拨号连接
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	tcpError(err, "net.Dial")
	// 读取键盘输入
	reader := bufio.NewReader(os.Stdin)
	// 创建缓冲区
	buffer := make([]byte, 1024)
	for {
		// 每次读取一行
		lineBytes, _, _ := reader.ReadLine()
		// 输入写入字符串
		conn.Write(lineBytes)
		// 读取缓冲区的数据(这块是服务端写的,这里就可以直接拿出来)
		n, _ := conn.Read(buffer)
		// 把字节转为字符串
		serverMsg := string(buffer[:n])
		fmt.Println("server", serverMsg)
		if serverMsg == "bye" {
			break
		}
	}
	fmt.Println("客户端以下线")

}
