package main

import "fmt"

func main() {
	//_17Testvar()

	_17testDefer()
}

func _17Testvar() {
	// x 已声明，y 没有声明
	var x int

	x, _ = _17f()
	// 这行编译出错，不能多次声明
	//x, _ := _17f()

	x, y := _17f()
	fmt.Println(x, y)
	x, y = _17f()
	fmt.Println(x, y)

	c := y
	fmt.Println(c)
}

func _17f() (x, y int) {
	return 0, 1
}

func _17increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func _17increaseB() (r int) {
	defer func() {
		r++
	}()

	return r
}

func _17testDefer() {
	fmt.Println(_17increaseA()) // 0
	fmt.Println(_17increaseB()) // 1
}
