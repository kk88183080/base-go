package main

import "fmt"

func main() {
	_21test()
	_20sfTest()
}

func _21test() {
	// 这个不会分配内存,优先选择
	var a []int
	// 这个会分配内存
	var b = []int{}

	fmt.Println(a, b)

	// 会出下标越界错误
	//a[0] = 1
	//fmt.Println(a, b)
	// 会出下标越界错误
	//b[0] = 1
	//fmt.Println(b)

	a = append(a, 1, 2)
	fmt.Println(a)
}

type _20S struct {
}

func _20f(x interface{}) {

}

// 这个方法不要这么定义
func _20g(x *interface{}) {

}

func _20testpara() {
	s := _20S{}
	p := &s

	_20f(s)
	//_20g(s)

	_20f(p)
	//_20g(p)
}

// 编译出错
// 永远不要使用一个指针指向一个接口类型，因为它已经是一个指针。

type _20SS struct {
	m string
}

func _20sf() *_20SS {
	return &_20SS{"ljw"}
}

func _20sfTest() {
	// 下面两种方式都可以
	p := _20sf()
	//p := *_20sf()
	fmt.Println(p.m)
}
