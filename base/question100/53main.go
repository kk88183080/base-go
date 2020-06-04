package main

import "fmt"

func main() {
	//_53test()
	//_53test2()
	_53test3()
}

type _53T struct {
	n int
}

func (y *_53T) Set(n int) {
	y.n = n
}

func _53getT() _53T {
	return _53T{}
}

//func _53test() {
//	_53getT().Set(1)
//}

// 编译就出错了
// _53T{}不可寻址，所以也就无法调用这个带结构体指针做为接受者的方法

// 修正方法1
func _53test2() {
	var t = _53getT()
	t.Set(1)
	fmt.Println(t.n)
}

// 修正方法2
func _53test3() {
	var t = _53getT()

	p := &t.n
	*p = 10

	fmt.Println(*p)
	fmt.Println(p)

}
