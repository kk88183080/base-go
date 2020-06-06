package main

import "fmt"

func main() {
	//_78test()
	_78test1()
}

func _78alwaysFalse() bool {
	return false
}

func _78test() {
	switch _78alwaysFalse(); {
	case true:
		fmt.Println(true)
	case false:
		fmt.Println(false)
	}
}

// 返回true Go 代码断行规则。

func _78test1() {
	switch _78alwaysFalse() {
	case true:
		fmt.Println(true)
	case false:
		fmt.Println(false)
	}
}

// 返回false
