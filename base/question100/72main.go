package main

import "fmt"

func main() {
	_72test2()
	_72test3()
}

type _72Orange struct {
	Qun int
}

func (o *_72Orange) incr(n int) {
	o.Qun += n
}

func (o *_72Orange) decr(n int) {
	o.Qun -= n
}

func (o *_72Orange) String() string {
	return fmt.Sprintf("String %#v", o.Qun)
}

func _72test2() {
	var org _72Orange
	org.incr(10)
	org.decr(5)
	fmt.Println(org) //String() 是指针方法，而不是值方法，所以使用 Println() 输出时不会调用到 String() 方法。
	fmt.Println(&org)
}

func _72test3() {
	org := &_72Orange{}
	org.incr(10)
	org.decr(5)
	fmt.Println(org)
}
