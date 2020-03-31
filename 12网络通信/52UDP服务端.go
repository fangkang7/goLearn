package main

import (
	"fmt"
	"net"
	"os"
)

func HandleError(err error, when string) {
	if err != nil {
		fmt.Println("错误地址,err", when, err)
		// 非正常退出,正常退出为0
		os.Exit(0)
	}
}

func main() {
	// 解析得到UDP地址 返回指针
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8888")
	HandleError(err, "ResolveUDPAddr")
	// 在UDP地址上建立UDP监听，得到连接  传入指针类型
	udpConn, err := net.ListenUDP("udp", udpAddr)
	HandleError(err, "ListenUDP")
	buffer := make([]byte, 1024)
	// 从连接中读入内容,丢入缓冲区
	n, remoteAddr, err := udpConn.ReadFromUDP(buffer)
	HandleError(err, "udpConn")
	fmt.Printf("读到来自%v的内容：%s", remoteAddr, string(buffer[:n]))
	udpConn.WriteToUDP([]byte("朕以阅"), remoteAddr)
}
