package main

import (
	"fmt"
	"go_code/family/utils"
)

func main() {
	fmt.Println("面向对象的记账功能")
	utils.NewFamily().ShowMenu()
}
