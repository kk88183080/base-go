package main

import "fmt"

func main() {
	_28test1()
}

func _28test() {
	/*
		// 编译出错-- 一个是切片，一个是数组，不同类型不能比较
		fmt.Println([...]int{1} == [2]int{1})
		// 切片不能
		fmt.Println([]int{1} == []int{1})
	*/
}

var _28p *int

func _28foo() (*int, error) {

	var i int = 5
	return &i, nil
}

func _28bar() {
	fmt.Println(*_28p)
}

func _28test1() {
	var err error
	_28p, err = _28foo()
	// 会导致编译出错
	//_28p, err := _28foo()
	if err != nil {
		fmt.Println(err)
		return
	}

	_28bar()
	fmt.Println(*_28p)
}
