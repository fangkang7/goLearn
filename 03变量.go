package main

/**
变量：存在于某个内存地址，值可以被修改(类型不可修改)
 */

 var kaka string = "我是咔咔";
 var (
 	age int = 13;
 	name string = "咔咔博客地址blog.fangkang.top";
 );

func main() {
	//age = 14;
	age = "咔咔博客地址blog.fangkang.top";
	print(age)
	// 声明和复制一站式搞定(仅限于函数内部使用)
	kakaBlog := "咔咔博客地址";
	print(kakaBlog)
}