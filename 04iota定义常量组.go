package main

import "fmt"

const (
	Sunday int = iota
	Monday
	Wednesday
	Thursday
	Tuesday
	Friday
	Saturday
)

func main() {
	fmt.Println(Sunday, Monday, Wednesday, Thursday, Tuesday, Friday, Saturday)
}
