package main

import "fmt"

func main() {
	//_52test()
	//_52test1()
	_52test2()
}

type _52X struct {
}

func (x *_52X) test() {
	fmt.Println(x)
}

/*func _52test() {
	var a *_52X
	a.test()
	_52X{}.test()
}*/

//X{} 是不可寻址的，不能直接调用方法。知识点：在方法中，指针类型的接收者必须是合法指针（包括 nil）,或能获取实例地址。
// 修复代码

func _52test1() {
	var a *_52X
	a.test()

	i2 := _52X{}
	i2.test()

	i3 := &_52X{}
	i3.test()
}

func _52test2() {

	var a = map[string]string{"one": "l", "two": "j", "three": "w"}
	//if v := a["one"]; v =="" {
	//	fmt.Println("one is not exits")
	//}

	if v, ok := a["one"]; ok {
		fmt.Printf("one is not exits, v:%s\n", v)
	}

	val := a["one"]
	fmt.Println(val)
}
