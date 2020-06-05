package main

import "fmt"

func main() {
	//_58test()
	_58test2()
}

type _58T struct {
	x int
	y *int
}

func _58test() {

	i := 10
	t := _58T{20, &i}

	p := &t.x // 是取的x指针地址

	*p++
	*p--

	t.y = p

	fmt.Println(*t.y)
}

//算符优先级。如下规则：递增运算符 ++ 和递减运算符 -- 的优先级低于解引用运算符 * 和取址运算符 &，解引用运算符和取址运算符的优先级低于选择器 . 中的属性选择操作符。

func _58test2() {
	x := make([]int, 2, 10)
	y := x[6:10] // len 10-6=4 cap 10-6=4
	fmt.Printf("6,10: len:%d, cap:%d\n", len(y), cap(y))

	// 下标越界
	//z := x[6:]   // len 2-6=-4 cap 10-6=4
	//fmt.Printf("6:, len:%d, cap:%d\n", len(z), cap(z))
	//截取符号 [i:j]，如果 j 省略，默认是原切片或者数组的长度，x 的长度是 2，小于起始下标 6 ，所以 panic。

	w := x[2:] // len 2-2=0 cap 10-2=8
	fmt.Printf("2:, len:%d, cap:%d\n", len(w), cap(w))
	// := 编译错误
	//_ := x[2:]
}
