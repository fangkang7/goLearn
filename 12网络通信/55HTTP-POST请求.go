package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func postError(err error, when string) {
	if err != nil {
		fmt.Println("错误地址,err", when, err)
		// 非正常退出,正常退出为0
		os.Exit(0)
	}
}
func main() {
	url := "https://www.fangkang.top/api/subjectByType"
	resp, err := http.Post(
		// 请求地址
		url,
		// 携带的数据类型
		"application/x-www-form-urlencoded",
		// 携带的参数
		strings.NewReader("st_id=1"))
	postError(err, "http.Post")
	// response属于io资源，用完需要关闭
	defer resp.Body.Close()
	// 获取body的内容
	bytes, err := ioutil.ReadAll(resp.Body)
	postError(err, "ioutil.ReadAll")
	// 返回的是字节需要转为字符串
	fmt.Println(string(bytes))
}
