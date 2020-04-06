package utils

import "fmt"

type familyAccount struct {
	// 声明一个变量,保存用户输入的值
	key string
	// 声明一个变量是否退出循环
	loop bool
	// 定义账户余额
	balance float64
	// 每次收支的金额
	money float64
	// 每次收支的说明
	note string
	// 收支的详情
	detaile string
	// 定义一个flag来判断是否有收支记录
	flag   bool
	choice string
}

/**
使用工厂模式返回一个指针
*/
func NewFamily() *familyAccount {
	return &familyAccount{
		// 声明一个变量,保存用户输入的值
		key: "",
		// 声明一个变量是否退出循环
		loop: true,
		// 定义账户余额
		balance: 1000.0,
		// 每次收支的金额
		money: 0.0,
		// 每次收支的说明
		note: "",
		// 收支的详情
		detaile: "收支\t 账户金额\t 收支余额\t 说    明",
		// 定义一个flag来判断是否有收支记录
		flag:   false,
		choice: "",
	}
}

/**
显示详情
*/
func (this *familyAccount) detail() {
	fmt.Println("-------------当前收支明细记录-----------------")
	if this.flag {
		fmt.Println(this.detaile)
	} else {
		fmt.Println("还没有任何记录,快去记录吧!")
	}
}

/**
收入
*/
func (this *familyAccount) income() {
	fmt.Println("本次收支金额:")
	fmt.Scan(&this.money)
	this.balance += this.money
	fmt.Println("本次收支说明")
	fmt.Scan(&this.note)
	this.detaile += fmt.Sprintf("\n收入 \t%v \t%v   \t%v\n", this.balance, this.money, this.note)
	this.flag = true
}

/**
支出
*/
func (this *familyAccount) expenditure() {
	fmt.Println("本次支出金额:")
	fmt.Scan(&this.money)
	// 对支出金额做一个判断
	if this.money > this.balance {
		fmt.Println("你的余额不够了,咔咔")
		return
	}
	// 从余额中减掉
	this.balance -= this.money
	fmt.Println("本次支出说明")
	fmt.Scan(&this.note)
	this.detaile += fmt.Sprintf("\n支出 \t%v \t%v   \t%v\n", this.balance, this.money, this.note)
	this.flag = true
}

/**
退出
*/
func (this *familyAccount) exit() {
	fmt.Println("你确定要退出吗? y/n")
	this.choice = ""
	for {
		fmt.Scan(&this.choice)
		if this.choice == "y" || this.choice == "n" {
			break
		}
		fmt.Println("你的输入有误,请重新输入 y/n")
	}
	if this.choice == "y" {
		this.loop = false
	}
}

/**
显示菜单
*/
func (this *familyAccount) ShowMenu() {
	for {
		fmt.Println("-------------家庭收支记账软件-----------------")
		fmt.Println("				 1.收支明细")
		fmt.Println("				 2.登记收入")
		fmt.Println("				 3.登记支出")
		fmt.Println("				 4.退出软件")
		fmt.Print("请选择1-4: ")

		fmt.Scan(&this.key)

		switch this.key {
		case "1":
			this.detail()
		case "2":
			this.income()
		case "3":
			this.expenditure()
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确1的选项...")
		}
		if this.loop == false {
			break
		}
	}
}
