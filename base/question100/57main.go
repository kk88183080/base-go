package main

import "fmt"

func main() {
	//_57test()
	_57test2()
}

func _57test() {
	var x interface{}

	var y interface{} = []int{1, 2, 3}

	a := x == x
	b := x == y

	fmt.Println(a)
	fmt.Println(b)

	// 执行出错
	//_ = y == y
	//因为两个比较值的动态类型 不是同一个类型

	// 编译出错
	//var z = []int{1}
	//_ = z == y

	// 执行出错 切片类型不可以比较
	var w interface{} = []int{2}
	_ = y == w

}

var _57o = fmt.Println

func _57test2() {
	a := [3]struct{}{}

	fmt.Printf("%T\n", a)

	c := make(chan int, 1)
	for range a {
		select {
		default:
			_57o(1)
		case <-c:
			_57o(2)
			c = nil
		case c <- 1:
			_57o(3)

		}
	}
}

//第一次循环，写操作已经准备好，执行 o(3)，输出 3；
//第二次，读操作准备好，执行 o(2)，输出 2 并将 c 赋值为 nil；
//第三次，由于 c 为 nil，走的是 default 分支，输出 1。
