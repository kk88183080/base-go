package main

import "fmt"

func main() {
	//_56test()
	_56test2()
}

func _56test() {

	s := make([]int, 3, 9)
	fmt.Println(len(s)) // len 3, cap 9
	fmt.Println(cap(s))

	s1 := s[4:8] // len j-i = 8-4 =4 cap 9-4=5
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
}

type _56N int

func (n _56N) _56Ntest() {
	fmt.Println(n)
}

func _56test2() {
	var n _56N = 10

	n++
	fmt.Println(n) // 11

	n++
	f1 := n._56Ntest // 12

	n++
	f2 := n._56Ntest // 13

	f1()           //12
	f2()           //13
	fmt.Println(n) //13

}

//方法值。当指针值赋值给变量或者作为函数参数传递时，会立即计算并复制该方法执行所需的接收者对象，与其绑定，以便在稍后执行时，能隐式第传入接收者参数。
