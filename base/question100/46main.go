package main

import "fmt"

func main() {
	//_46test()
	_46test1()
}

func _46test() {
	const a = 0
	const b = 0
}

// 常量是一个简单值的标识符，在程序运行时，不会被修改的量。不像变量，常量未使用是能编译通过的。

const (
	_46x uint16 = 120
	_46y
	_46s = "ljw"
	_46z
)

func _46test1() {
	fmt.Printf("%T %v\n", _46y, _46y)
	fmt.Printf("%T %v\n", _46z, _46z)
}

//常量组中如不指定类型和初始化值，则与上一行非空常量右值相同

/*
func _46test2() {
	var x string = nil
	if x == nil {
		fmt.Println("default")
	}
}*/
//将 nil 分配给 string 类型的变量。这是个大多数新手会犯的错误
// 不能给string 分配 nil值

// 修复如下
func _46test2() {
	var x string
	if x == "" {
		fmt.Println("default")
	}
}
