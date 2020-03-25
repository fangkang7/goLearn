package main

import "testing"

func BenchmarkSaveHuman(b *testing.B) {
	b.Log("测试开始")
	// 测试开销
	b.ReportAllocs()
	human := &Human{
		Name: "咔咔",
		Age:  24,
	}
	for i := 0; i < b.N; i++ {
		SaveHuman(human, "C:/Users/Administrator/Desktop/go/11压力测试/kaka.json")
	}
}

func BenchmarkLoadHuman(b *testing.B) {
	b.Log("测试开始")
	// 测试开销
	b.ReportAllocs()
	hPrt := new(Human)
	for i := 0; i < b.N; i++ {
		LoadHuman("C:/Users/Administrator/Desktop/go/11压力测试/kaka.json", hPrt)
	}
}
