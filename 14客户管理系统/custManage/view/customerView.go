package main

import "fmt"

/**
这个结构体不用大写，是因为这个结构体就在这个文件使用
 */
type customerView struct {
	// 接收用户的输入的值
	key string
	// 表示是否循环的显示主菜单
	loop bool
}

/**
显示主菜单
 */
func (this *customerView)mainMenu()  {
	// 死循环显示客户信息管理系统
	for  {
		fmt.Println("------------客户信息管理系统-------------")
		fmt.Println("             1 添 加 客 户")
		fmt.Println("             2 修 改 客 户")
		fmt.Println("             3 删 除 客 户")
		fmt.Println("             4 客 户 列 表")
		fmt.Println("             5 退      出")

		fmt.Println("请输入1-5")
		// 这里是获取用户的输入的值
		fmt.Scan(&this.key)

		switch this.key {
		case "1":
			fmt.Println("添 加 客 户")
		case "2":
			fmt.Println("修 改 客 户")
		case "3":
			fmt.Println("删 除 客 户")
		case "4":
			fmt.Println("客 户 列 表")
		case "5":
			this.loop= false
		default:
			fmt.Println("您输入的有误,请重新输入")
		}
		// 如果用户输入的是5那么就是退出客户系统
		if !this.loop{
			break
		}
	}
	fmt.Println("您退出了客户关系系统")
}

func main() {
	// 创建一个customer的实例，并显示主菜单
	customerView := customerView{
		key:"",
		loop:true,
	}
	customerView.mainMenu()
}
