package main

import "fmt"

func main() {
	_22deferTest()
}

// 两个地方有语法问题。golang 的字符串类型是不能赋值 nil 的，也不能跟 nil 比较。
/*
func _22strTest()  {
	var s string=nil
	if s == nil {
		fmt.Println("s is nil")
	}

	fmt.Println(s)
}
*/

var flag bool = true

func _22deferTest() {

	defer func() {
		fmt.Println("1")
	}()

	if flag {
		fmt.Println("2")
		return
	}

	defer func() {
		fmt.Println("3")
	}()
}
