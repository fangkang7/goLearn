package main

import (
	"fmt"
	"reflect"
)

type PeopleParent struct {
	Kaka string
}

type People struct {
	PeopleParent
	Name string
	Age  int
}

func (p People) Eat(name string) {
	fmt.Println("咔咔在吃什么呢！", name)
	p.Name = name
}

func main() {
	p := People{
		PeopleParent: PeopleParent{Kaka: "咔咔的父类属性"},
		Name:         "咔咔",
		Age:          24,
	}
	typeAPI(p)
	valueAPI(p)
}
func typeAPI(obj interface{}) {
	// 返回保存值的类型
	oType := reflect.TypeOf(obj)
	fmt.Println(oType) // main.People
	// 原始类型
	kind := oType.Kind()
	fmt.Println(kind) // struct
	// 类型名称
	fmt.Println(oType.Name()) // People
	// 属性和方法的个数
	fmt.Println(oType.NumField())  // 2
	fmt.Println(oType.NumMethod()) // 0
	// 获取全部属性
	for i := 0; i < oType.NumField(); i++ {
		structField := oType.Field(i)
		// name string    age int
		fmt.Println(structField.Name, structField.Type)
	}
	// 获取全部方法
	for i := 0; i < oType.NumMethod(); i++ {
		structMethod := oType.Method(i)
		fmt.Println(structMethod.Name, structMethod.Type)
	}
	// 获取父类的属性  []int{0, 0}获取第0个父类  第0个属性
	fmt.Println(oType.FieldByIndex([]int{0, 0}).Name)
}

func valueAPI(p People) {
	valueOf := reflect.ValueOf(p)
	//valueOf := reflect.ValueOf(&p)
	//byEat := valueOf.MethodByName("Eat")
	//byEat.Call([]reflect.Value{reflect.ValueOf("西瓜")})
	//fmt.Println(valueOf)
	// 获取所有属性值
	for i := 0; i < valueOf.NumField(); i++ {
		value := valueOf.Field(i)
		// {}
		//咔咔
		//24
		fmt.Println(value)
	}
	// 获取父类属性
	fieldByIndex := valueOf.FieldByIndex([]int{0, 0})
	fmt.Println(fieldByIndex.Interface()) // 咔咔的父类属性

	// 获得指针value的内容进而获得成员的值
	valuePrt := reflect.ValueOf(&p)
	elem := valuePrt.Elem()
	value := elem.Field(0).Interface()
	fmt.Println(value) //{咔咔的父类属性}

	// 根据属性名获取值
	age := elem.FieldByName("Age")
	fmt.Println("咔咔的年龄", age) // 咔咔的年龄 24
	// 修改属性值
	elem.FieldByName("Age").SetInt(26)
	fmt.Println(elem) //{{咔咔的父类属性} 咔咔 26}

	// 调用对象的方法
	mValue := elem.MethodByName("Eat")
	// 参数需要反射
	mValue.Call([]reflect.Value{reflect.ValueOf("西瓜")})
	fmt.Println(elem) //咔咔在吃什么呢！ 西瓜
}
