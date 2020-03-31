package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func getError(err error, when string) {
	if err != nil {
		fmt.Println("错误地址,err", when, err)
		// 非正常退出,正常退出为0
		os.Exit(0)
	}
}
func main() {
	url := "https://fangkang.top"
	resp, err := http.Get(url)
	getError(err, "http.Get")
	// response属于io资源，用完需要关闭
	defer resp.Body.Close()
	// 获取body的内容
	bytes, err := ioutil.ReadAll(resp.Body)
	getError(err, "ioutil.ReadAll")
	// 返回的是字节需要转为字符串
	fmt.Println(string(bytes))
}
