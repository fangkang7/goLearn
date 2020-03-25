package main

import (
	"encoding/json"
	"os"
)

func SaveHuman(human *Human, filename string) (bool, error) {
	// 打开文件
	file, _ := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	// 挂起文件最后执行关闭
	defer file.Close()
	// 创建解码器
	encoder := json.NewEncoder(file)
	err := encoder.Encode(human)
	return err == nil, err
}

func LoadHuman(filename string, human *Human) error {
	file, _ := os.OpenFile(filename, os.O_RDONLY, 0666)
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(human)
	return err
}
