package main

import "fmt"

func main() {
	//_44test()
	//_44test2()
	_44test4()
}

func _44test() {
	m := make(map[string]int, 3)
	//fmt.Println(cap(m))
	fmt.Println(len(m))
}

//使用 cap() 获取 map 的容量。
//1.使用 make 创建 map 变量时可以指定第二个参数，不过会被忽略。
//2.cap() 函数适用于数组、数组指针、slice 和 channel，不适用于 map，可以使用 len() 返回 map 的元素个数。

func _44test2() {
	//var x =nil
	var x func() = nil
	_ = x
}

// nil 用于表示 interface、函数、maps、slices 和 channels 的“零值”。
// 如果不指定变量的类型，编译器猜不出变量的具体类型，导致编译错误。

type _44info struct {
	result int
}

func _44work() (int, error) {
	return 13, nil
}

func _44test4() {
	var data _44info
	data.result, _ = _44work()

	// 不能使用短变量声明设置结构体字段值
	//data.result, err := _44work()
	fmt.Println(data)
}
