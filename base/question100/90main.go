package main

import "fmt"

func main() {
	//_90test()
	_90test1()
}

type _90T int

func _90F(t _90T) {
	fmt.Println(t)
}

func _90test() {
	var q int
	_90F(_90T(q))
	// 编译出错
	//_90F(q)
}

type _90T1 []int

func _90F1(n []int) {
	fmt.Println(n)
}

func _90test1() {
	var q []int = []int{1, 2}
	q[0] = 1
	_90F1(q)
}
