package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 定义工人接口
type Worker interface {
	// 每天工作多少小时，产出何种产品
	Work(hour int) (chanpin string)
	// 休息
	Rest()
}

// 定义码农
type Coder struct {
	skill string
}

func (Coder *Coder) Work(hour int) (chanpin string) {
	fmt.Printf("码农一天工作%d小时\n", hour)
	return "BUG"
}

func (Coder *Coder) Rest() {
	fmt.Println("休息是什么？？？")
}

// 定义产品经理
type ProductManager struct {
	skill string
}

func (P *ProductManager) Work(hour int) (chanpin string) {
	fmt.Printf("产品一天工作%d小时\n", hour)
	return "无逻辑的需求"
}

func (P *ProductManager) Rest() {
	fmt.Println("产品可以使劲的休息")
}

// 定义boos
type Boos struct {
	skill string
}

func (Boos *Boos) Work(hour int) (chanpin string) {
	fmt.Printf("boos一天工作%d小时\n", hour)
	return "梦想"
}

func (Boos *Boos) Rest() {
	fmt.Println("无休息")
}

func main() {
	// 创建一个工人切片保存三种职业
	// 这里需要注意一个点  这个切片的名字Worker需要跟接口名一致
	workers := make([]Worker, 0)

	// 往切片添加的都是指针并非值，因为在方法主语用的是指针形式 (Boos *Boos)
	workers = append(workers, &Coder{skill: "写代码"})
	workers = append(workers, &ProductManager{skill: "提需求"})
	workers = append(workers, &Boos{skill: "想方案"})
	// 创建一个种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 创建一个随机数代表星期几
	weekday := r.Intn(7)

	if weekday > 0 && weekday < 6 {
		for _, Worker := range workers {
			Worker.Work(8)
		}
	} else {
		for _, Worker := range workers {
			Worker.Rest()
		}
	}
}
