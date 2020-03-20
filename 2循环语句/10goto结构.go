package main

import "fmt"

func main() {
GameStart:
	fmt.Println("开业大吉")
	goto GameOver
	fmt.Println("此处有红包")
GameOver:
	fmt.Println("红包已经领完，请明天再来")
	for i := 0; i < 10; i++ {
		fmt.Println("咔咔欢迎你的到来")
	}
	goto GameStart
}
