package main

import (
	"fmt"
	"os"
)

func main() {
	//_88test()
	//_88test2()
	//_88test3()
	_88test4()
}

func _88test() {

	m := make(map[string]int)
	m["foo"]++
	fmt.Println(m["foo"])
}

func _88test2() {
	m := make(map[string]int)
	v := m["foo"]
	v++
	m["foo"] = v
	//v++
	fmt.Println(m["foo"])
}

func _88Foo() error {
	var err *os.PathError = nil
	return err
}

func _88test3() {
	err := _88Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
	fmt.Printf("%#v\n", err)
	fmt.Printf("%#v\n", nil)
}

//接口值与 nil 值。只有在值和动态类型都为 nil 的情况下，接口值才为 nil。
//Foo() 函数返回的 err 变量，值为 nil、动态类型为 *os.PathError，
//与 nil（值为 nil，动态类型为 nil）显然是不相等。我们可以打印下变量 err 的详情：

func _88foo1() error {
	return nil
}

func _88test4() {
	err := _88foo1()
	fmt.Println(err)
	fmt.Println(err == nil)
}
