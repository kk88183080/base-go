package main

import "fmt"

func main() {
	nil := 123
	fmt.Println(nil)
	//var _ map[string]int = nil
	// 编译出错，类型不对

	_62test1()
	_62test2()
}

func _62test() {
	_, err := fmt.Println("x")
	if err != nil {
		fmt.Println("出错了！！！")
	}
}

func _62test1() {
	var x int8 = -128
	var y = x / -1

	fmt.Println(y)
}

// 因为溢出 int8的最大值是-128~127 -128/-1=128 128大于最大值127 相等于是127+1=-128

func _62test2() {
	var x int8 = 127
	fmt.Println(x + 1)
}
