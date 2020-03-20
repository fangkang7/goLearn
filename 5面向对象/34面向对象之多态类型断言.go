package main

import (
	"fmt"
)

// 定义工人接口
type Worker1 interface {
	// 每天工作多少小时，产出何种产品
	Work1(hour int) (chanpin string)
	// 休息
	Rest1()
}

// 定义码农
type Coder1 struct {
	skill string
}

func (Coder *Coder1) Work1(hour int) (chanpin string) {
	fmt.Printf("码农一天工作%d小时\n", hour)
	return "BUG"
}

func (Coder *Coder1) Rest1() {
	fmt.Println("休息是什么？？？")
}

// 定义产品经理
type ProductManager1 struct {
	skill string
}

func (P *ProductManager1) Work1(hour int) (chanpin string) {
	fmt.Printf("产品一天工作%d小时\n", hour)
	return "无逻辑的需求"
}

func (P *ProductManager1) Rest1() {
	fmt.Println("产品可以使劲的休息")
}

// 定义boos
type Boos1 struct {
	skill string
}

func (Boos *Boos1) Work1(hour int) (chanpin string) {
	fmt.Printf("boos一天工作%d小时\n", hour)
	return "梦想"
}

func (Boos *Boos1) Rest1() {
	fmt.Println("无休息")
}

func main() {
	// 创建一个工人切片保存三种职业
	// 这里需要注意一个点  这个切片的名字Worker需要跟接口名一致
	workers := make([]Worker1, 0)

	// 往切片添加的都是指针并非值，因为在方法主语用的是指针形式 (Boos *Boos)
	workers = append(workers, &Coder1{skill: "写代码"})
	workers = append(workers, &ProductManager1{skill: "提需求"})
	workers = append(workers, &Boos1{skill: "想方案"})
	// 创建一个种子
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 创建一个随机数代表星期几
	//weekday := r.Intn(7)

	// 类型断言方式一
	/**
	休息是什么？？？
	任何指针都不属于
	无休息
	*/
	for _, s := range workers {
		switch s.(type) {
		case *Coder1:
			s.Rest1()
		case *Boos1:
			s.Rest1()
		default:
			fmt.Println("任何指针都不属于")
		}
	}

	// 类型断言方式二
	/**
	  如果是:ok为true，Coder1Ptr有值
	  如果不是：ok为false，Coder1Ptr为nil
	*/
	for _, s := range workers {
		if Coder1Ptr, ok := s.(*Coder1); ok {
			// 码农一天工作10小时
			Coder1Ptr.Work1(10)
		}
	}

}
