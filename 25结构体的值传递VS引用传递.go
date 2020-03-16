package main

import "fmt"

type Person struct {
	name string
	age  int
	kaka []string
}

func main() {
	p := Person{name: "咔咔", age: 24, kaka: []string{"dalanqiu"}}
	// 值传递，传入的是并非是咔咔本人，而是咔咔的副本
	// 对副本的修改并不能影响咔咔本人
	//showPerson1(p)

	// 传入的是结构的本身(地址)，并非副本
	showPerson2(&p)
	// showPerson2方法修改的值
	fmt.Println(p.name)
	// p的地址是0x11019a20
	fmt.Printf("p的地址是%p\n", &p)
}

func showPerson1(p Person) {
	// p的地址是0x11019a20
	fmt.Printf("p的地址是%p\n", &p)
	// 咔咔
	fmt.Println(p.name)
	// 24
	fmt.Println(p.age)
	// [dalanqiu]
	fmt.Println(p.kaka)
	p.name = "我不是咔咔"
}

/**
这里的参数是指针类型 在传值的时候就需要指针类型
*/
func showPerson2(p *Person) {
	// p的地址是0x11019a20
	fmt.Printf("p的地址是%p\n", &p)
	// 咔咔
	fmt.Println(p.name)
	// 24
	fmt.Println(p.age)
	// [dalanqiu]
	fmt.Println(p.kaka)
	p.name = "showPerson2方法修改的值"
}
