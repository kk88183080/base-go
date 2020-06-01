package main

import "fmt"

func main() {
	_37test1()
}

// 都正确
//A. 给一个 nil channel 发送数据，造成永远阻塞
//B. 从一个 nil channel 接收数据，造成永远阻塞
//C. 给一个已经关闭的 channel 发送数据，引起 panic
//D. 从一个已经关闭的 channel 接收数据，如果缓冲区中为空，则返回一个零值

const i = 100

var j = 123

func _37test() {
	fmt.Println(&j, j)
	// 编译出错
	// 常量不同于变量的在运行期分配内存，常量通常会被编译器在预处理阶段直接展开，作为指令数据使用，所以常量无法寻址。
	//fmt.Println(&i, i)
}

func _37GetValue(m map[int]string, id int) (string, bool) {
	if s, exists := m[id]; exists {
		fmt.Println(s, exists)
		return "exists", true
	}

	// 编译出错
	// 函数返回值类型。nil 可以用作 interface、function、pointer、map、slice 和 channel 的“空值”。
	// 但是如果不特别指定的话，Go 语言不能识别类型，所以会报错:cannot use nil as type string in return argument.
	//return nil, false
	return "", false
}

func _37test1() {
	intMap := map[int]string{
		1: "l",
		2: "s",
	}

	v, err := _37GetValue(intMap, 1)
	fmt.Println(v, err)
}
