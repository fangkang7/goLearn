package model

/**
声明一个Customer的结构体，表示客户信息
 */
type Customer struct {
	Id int
	// 名字
	Name string
	// 性别
	Gender string
	// 年龄
	Age int
	// 手机号
	Phone string
	// 邮箱
	Email string
}

/**
写一个工厂模式，返回一个Customer的实例
 */
func NewCustomer(id int,name string,gender string,age int,phone string,email string) Customer {
	return Customer{
		Id:id,
		Name:name,
		Gender:gender,
		Age:age,
		Phone:phone,
		Email:email,
	}
}
