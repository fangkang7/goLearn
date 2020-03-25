package main

import "testing"

func TestSaveHuman(t *testing.T) {
	human := &Human{
		Name: "咔咔1",
		Age:  24,
	}
	ok, err := SaveHuman(human, "C:/Users/Administrator/Desktop/go/10单元测试/kaka1.json")
	if !ok {
		t.Error("TestSaveHuman测试失败", err)
		return
	}
	LoadHumans := new(Human)
	err = LoadHuman("C:/Users/Administrator/Desktop/go/10单元测试/kaka1.json", LoadHumans)
	if err != nil {
		t.Error("TestSaveHuman 测试失败")
	}
	if LoadHumans.Name == human.Name && LoadHumans.Age == human.Age {
		t.Log("TestSaveHuman测试成功")
	} else {
		t.Error("测试失败")
	}
}

func TestLoadHuman(t *testing.T) {
	human := &Human{
		Name: "咔咔2",
		Age:  12,
	}
	ok, err := SaveHuman(human, "C:/Users/Administrator/Desktop/go/10单元测试/kaka2.json")
	if !ok {
		t.Error("TestLoadHuman测试失败", err)
		return
	}
	LoadHumans := new(Human)
	err = LoadHuman("C:/Users/Administrator/Desktop/go/10单元测试/kaka2.json", LoadHumans)
	if err != nil {
		t.Error("TestLoadHuman 测试失败")
	}
	if LoadHumans.Name == human.Name && LoadHumans.Age == human.Age {
		t.Log("TestLoadHuman测试成功")
	} else {
		t.Error("测试失败")
	}
}
