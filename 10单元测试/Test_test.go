package main

import "testing"

/**
测试斐波那契数列
*/
func TestGetFibonacci(t *testing.T) {
	fibonacci := GetFibonacci(6)
	if fibonacci != 13 {
		t.Errorf("测试失败")
		return
	}
	t.Logf("成功")
}

func TestGetRecursion(t *testing.T) {
	mMap := make(map[int]int)
	mMap[10] = 55
	mMap[100] = 5050
	for key, value := range mMap {
		if GetRecursion(key) != value {
			t.Errorf("测试失败")
			return
		}
	}
	t.Logf("测试成功  ")
}
