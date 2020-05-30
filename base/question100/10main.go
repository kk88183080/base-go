package main

import "fmt"

func main() {

	a := 1
	b := 3.4
	// 类型不一样，编译错误
	//c:=a/b
	//c := a + b
	//c := a - b
	//c := a * b

	// 先做类型转换，再做运算
	c := float64(a) + b
	fmt.Println(c)
}
