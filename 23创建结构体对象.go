package main

import "fmt"

type Person struct {
	name  string
	age   int
	aiaho []string
}

func main() {
	// 创建对象时，给所有属性依次赋值  切记顺序不可乱(这种方式是没有写键值  直接根据结构体写值所以必须按照顺序)
	p := Person{"kaka", 24, []string{"篮球", "写代码"}}
	// p的值main.Person{name:"kaka", age:24, aiaho:[]string{"篮球", "写代码"}}
	fmt.Printf("p的值%#v\n", p)

	// 有选择的创建结构体对象(有选择的创建结构体对象，这里把需要的属性已经提前写出来了)
	p = Person{aiaho: []string{"打篮球", "写代码"}, name: "咔咔"}
	// p的值main.Person{name:"咔咔", age:0, aiaho:[]string{"打篮球", "写代码"}}
	fmt.Printf("p的值%#v\n", p)

	// 创建一个main.Person的指针
	pPrt := new(Person)
	// pPrt的类型是*main.Person,值是&{ 0 []}
	fmt.Printf("pPrt的类型是%T,值是%v\n", pPrt, pPrt)
	pPrt.name = "咔咔手赚网：fangkang.top"
	pPrt.age = 24
	pPrt.aiaho = []string{"咔咔手赚网", "咔咔博客"}
	// pPrt的值是&{咔咔手赚网：fangkang.top 24 [咔咔手赚网 咔咔博客]}
	fmt.Printf("pPrt的值是%v\n", pPrt)

	// 既然pPtr是一个指针，那么我们就可以通过指针来获取值
	//pPer的值是{咔咔手赚网：fangkang.top 24 [咔咔手赚网 咔咔博客]}
	fmt.Printf("pPer的值是%v\n", *pPrt)
	// 取出pPer的地址  pPer的地址是0x11006120
	fmt.Printf("pPer的地址是%p\n", &pPrt)
}
