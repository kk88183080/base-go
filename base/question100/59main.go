package main

import "fmt"

func main() {
	//_59test()
	_59test2()
}

type _59N int

func (n *_59N) _59Ntest() {
	fmt.Println(*n)
}

func _59test() {
	var n _59N = 10
	p := &n

	n++
	f1 := n._59Ntest // 11

	n++
	f2 := p._59Ntest // 12

	n++
	fmt.Println(n) // 13

	f1() // 13
	f2() // 13
}

// 13 13 13
// 当目标方法的接收者是指针类型时，那么被复制的就是指针。

func _59test2() {
	var m map[int]bool
	_ = m[123]
	var p *[5]string

	for range p {
		_ = len(p)
		fmt.Println(len(p))
	}

	var s []int
	_ = s[:]
	s, s[0] = []int{1, 2}, 9
}

//因为左侧的 s[0] 中的 s 为 nil。
