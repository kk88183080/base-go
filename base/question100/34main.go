package main

import "fmt"

func main() {
	//_34test()
	_34test1()
}

type _34Myint int

func _34test() {
	var i int = 1
	// 编译出错
	//var j _34Myint = i.(_34Myint)
	var j _34Myint = _34Myint(i)
	fmt.Println(j)
}

type Integer int

func (o Integer) add(i Integer) Integer {
	return o + i
}

func (o *Integer) add1(i Integer) Integer {
	return *o + i
}

func _34test1() {
	var a Integer = 1
	var b Integer = 2
	var i interface{} = &a
	sum := i.(*Integer).add(b)
	fmt.Println(sum)

	sum = i.(*Integer).add1(b)
	fmt.Println(sum)
}
