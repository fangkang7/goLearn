package main

import "fmt"

func main() {
	// 声明一个变量,保存用户输入的值
	key := ""
	// 声明一个变量是否退出循环
	loop := true

	// 定义账户余额
	balance := 1000.0
	// 每次收支的金额
	money := 0.0
	// 每次收支的说明
	note := ""
	// 收支的详情
	detaile := "收支\t 账户金额\t 收支余额\t 说    明"
	// 定义一个flag来判断是否有收支记录
	flag := false
	// 显示这个主菜单
	for {
		fmt.Println("-------------家庭收支记账软件-----------------")
		fmt.Println("				 1.收支明细")
		fmt.Println("				 2.登记收入")
		fmt.Println("				 3.登记支出")
		fmt.Println("				 4.退出软件")
		fmt.Print("请选择1-4: ")

		fmt.Scan(&key)

		switch key {
		case "1":
			fmt.Println("-------------当前收支明细记录-----------------")
			if flag {
				fmt.Println(detaile)
			} else {
				fmt.Println("还没有任何记录,快去记录吧!")
			}
		case "2":
			fmt.Println("本次收支金额:")
			fmt.Scan(&money)
			balance += money
			fmt.Println("本次收支说明")
			fmt.Scan(&note)
			detaile += fmt.Sprintf("\n收入 \t%v \t%v   \t%v\n", balance, money, note)
			flag = true
		case "3":
			fmt.Println("本次支出金额:")
			fmt.Scan(&money)
			// 对支出金额做一个判断
			if money > balance {
				fmt.Println("你的余额不够了,咔咔")
				break
			}
			// 从余额中减掉
			balance -= money
			fmt.Println("本次支出说明")
			fmt.Scan(&note)
			detaile += fmt.Sprintf("\n支出 \t%v \t%v   \t%v\n", balance, money, note)
			flag = true
		case "4":
			fmt.Println("你确定要退出吗? y/n")
			choice := ""
			for {
				fmt.Scan(&choice)
				if choice == "y" || choice == "n" {
					break
				}
				fmt.Println("你的输入有误,请重新输入 y/n")
			}
			if choice == "y" {
				loop = false
			}
		default:
			fmt.Println("请输入正确1的选项...")
		}
		if loop == false {
			break
		}
	}
	fmt.Println("你退出了记账软件")
}
