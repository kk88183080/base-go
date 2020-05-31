package main

import "fmt"

func main() {
	//_15test()
	//_15test1()

	_15test3()
}

func _15test() {
	var a []int
	var b = []int{}

	if a == nil {
		fmt.Println("a is nil")
	} else {
		fmt.Println("a is not nil")
	}

	if b == nil {
		fmt.Println("b is nil")
	} else {
		fmt.Println("b is not nil")
	}
}

func _15test1() {
	i := 65
	fmt.Println(string(i))
}

// 执行方法
type _15A interface {
	ShowA() int
}

type _15B interface {
	ShowB() int
}

type _15Work struct {
	work int
}

func (w _15Work) ShowA() int {
	// 如查加上这两行，代码返回值就会有变化
	w.work += 10
	return w.work
	return w.work + 10
}

func (w _15Work) ShowB() int {
	return w.work + 20
}

func _15test3() {
	w := _15Work{work: 3}
	var a _15A = w
	var b _15B = w
	fmt.Println(a.ShowA())
	fmt.Println(b.ShowB())
}
