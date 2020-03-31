package main

import (
	"net/http"
)

func main() {
	// /hello是路由    writer响应的写器         request用户的请求
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("我是GO搭建的服务器"))
	})
	http.ListenAndServe("127.0.0.1:8080", nil)
}
