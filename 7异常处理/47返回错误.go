package main

import "fmt"

type Person struct {
}

// 这是一个结构体接口，这里直接使用实现即可
func (p *Person) Error() {
	fmt.Println("我是自定义的错误类型")
}

func main() {
	fmt.Println("请输入账号密码")
	var number int
	fmt.Scan(&number)
	if number == 1 {
		fmt.Println("账号正确")
	} else {
		person := new(Person)
		person.Error()
	}
}
