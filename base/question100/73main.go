package main

import "fmt"

func main() {
	_73test2()
}

func _73test() []func() {
	var funs []func()

	for i := 0; i < 2; i++ {
		funs = append(funs, func() {
			fmt.Println(&i, i)
		})
	}

	return funs
}

func _73test2() {
	funs := _73test()

	for _, f := range funs {
		f()
	}
}

// 闭包延迟求值。
// for 循环局部变量 i，匿名函数每一次使用的都是同一个变量。（说明：i 的地址，输出可能与上面的不一样）。
