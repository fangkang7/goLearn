package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

/**
所有商品的都有共性:名字,价钱
自行封装三种商品
随意给一个商品的切片,将每件商品的所有属性信息输出到json文件
使用反射实现
*/

/**
定义的父类  名字  价钱
*/
//type Product struct {
//	Name  string
//	price float64
//}

type Computer struct {
	Name  string
	Price float64
	Cpu   string
	// 内存
	Memory int
	// 磁盘
	Disk int
}

type TShirt struct {
	Name  string
	Price float64
	Color string
	Size  int
	Sex   bool
}

type Car struct {
	Name  string
	Price float64
	// 容量
	Cap   int
	Power string
}

func main() {
	// 定义一个可以存储任意类型的空切片
	products := make([]interface{}, 0)
	products = append(products, Computer{"未来人类", 100, "英特尔", 16, 100})
	products = append(products, TShirt{"衣服", 1000, "红色", 16, true})
	products = append(products, Car{"车子", 100, 6, "油电混动"})
	for _, p := range products {
		// 获取p对象的type和value
		pType := reflect.TypeOf(p)
		pValue := reflect.ValueOf(p)
		// 构造属性名和属性值的字典
		fileNameValueMap := make(map[string]interface{})
		// 从type中拿到所有属性名
		for i := 0; i < pType.NumField(); i++ {
			// 获取到所有的属性名
			fieldName := pType.Field(i).Name
			// 全部复制为nil
			fileNameValueMap[fieldName] = nil
		}
		// 从value中根据属性名拿到所有的属性值
		for fieldName, _ := range fileNameValueMap {
			// 返回的是一个反射值，我们需要使用interface来转为正射
			// 返回的是属性的值
			filedValue := pValue.FieldByName(fieldName).Interface()
			// 然后把属性名和属性值对应上
			fileNameValueMap[fieldName] = filedValue
		}
		fileName := "C:/Users/Administrator/Desktop/go/反射/" + fileNameValueMap["Name"].(string) + ".json"
		err := EncodeMapToJsonFile(fileNameValueMap, fileName)
		if err != nil {
			fmt.Println("写出失败")
		} else {
			fmt.Println("写出成功")
		}
	}
}

/**
写入json文件
*/
func EncodeMapToJsonFile(dataMpa map[string]interface{}, filename string) error {
	file, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 666)
	defer file.Close()
	encoder := json.NewEncoder(file)
	err := encoder.Encode(dataMpa)
	return err
}
