package main

import "fmt"

func main01() {
	var blogAddress string
	fmt.Println("请输入咔咔的博客地址")
	fmt.Scan(&blogAddress)

	if blogAddress == "blog.fangkang.top" {
		fmt.Println("是真的")
	} else {
		fmt.Println("是假的")
	}

	var kakaZhuan string
	fmt.Println("请输入咔咔手赚网的地址")
	fmt.Scan(&kakaZhuan)

	if kakaZhuan == "fangkang.top" {
		fmt.Println("咔咔手赚网的地址是对的")
	} else if kakaZhuan == "www.fangkang.top" {
		fmt.Println("这个地址也可以进入")
	} else {
		fmt.Println("你输入的是什么东东")
	}

}

func main() {
	var kakaZhuan string
	fmt.Println("请输入关于咔咔的所有网站")
	fmt.Scan(&kakaZhuan)

	/**
	关于switch也可以做条件选择根据项目自己定义即可
	*/
	switch kakaZhuan {
	case "fangkang.top":
		fmt.Println("这是咔咔手赚网地址")
	case "blog.fangkang.top":
		fmt.Println("这是咔咔博客地址")
	default:
		fmt.Println("咔咔的其他的网站正在开发中")
	}

}
