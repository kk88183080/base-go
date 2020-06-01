package main

import "fmt"

func main() {
	//_41test2()
	_41test3()
}

// 编码出错，{不能单独占一行
//func _41test()
//{
//	fmt.Println("xx")
//}

func _41test2() {
	var x = []int{2: 2, 3, 0: 1}

	fmt.Println(x) // 1 0 2 3
}

//输出[1 0 2 3]，字面量初始化切片时候，可以指定索引，没有指定索引的元素会在前一个索引基础之上加一，所以输出[1 0 2 3]，而不是[1 3 2]。

func _41incr(p *int) int {
	*p++
	return *p
}

func _41test3() {
	v := 1
	_41incr(&v)
	fmt.Println(v)
}

//指针。p 是指针变量，指向变量 v，*p++操作的意思是取出变量 v 的值并执行加一操作，所以 v 的最终值是 2。
