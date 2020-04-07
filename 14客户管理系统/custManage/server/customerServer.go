package server

import "go_code/custManage/model"

/**
该customerServer完成对model/Customer的操作，包括增删改查
 */
type Customer struct {
	// 创建一个结构体的切片
	customers []model.Customer
	// 声明一个字段，表示当前切片含有多少个客户(作为新客户的ID号)
	customerNum int
}
