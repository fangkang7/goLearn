package main

import "fmt"

/**
首先我们需要明确一个概念那就是什么是面向对象
面向对象简称OOP，是一种计算机编程架构，使用面向对象编程可以实现代码的重用性，灵活性，扩展性。
面向对象符合人类看待事物的一般规律。使其代码更简洁，更容易维护，并且具有更强的重用性

在PHP中面向对象的三大特性：封装、继承、多态
当然在go中也是一样的

在PHP中我们有类  但是在go使用的是结构体
*/

type Person struct {
	name string
	age  int
	like []string
}

func main() {
	// 创建对象并通过对象的指针去访问对象的属性和方法
	person := &Person{name: "咔咔", age: 24}
	// {咔咔 24 []}
	fmt.Println(*person)
	// 咔咔
	fmt.Println(person.name)
	// 24
	fmt.Println(person.age)
	// 咔咔博客地址：blog.fangkang.top
	person.eat()
	// 咔咔手赚网地址：fangkang.top
	person.play()
}

/**
无论方法的主语定义为指针还是值都是可以正常访问的
但是一般情况下我们都会把方法的主语定义为指针  因为传值的话是一个副本  指针是一个地址  这样在方法里边改变了属性的值 那么是通过这个地址去改变的
*/
func (p *Person) eat() {
	fmt.Println("咔咔博客地址：blog.fangkang.top")
}

func (p *Person) play() {
	fmt.Println("咔咔手赚网地址：fangkang.top")
}
