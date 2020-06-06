package main

import (
	"errors"
	"fmt"
)

func main() {
	//_64test()
	_64test2()
}

func _64test() {
	defer func() {
		rs := recover()
		fmt.Printf("error:%T, v:%v", rs, rs)
	}()

	//panic(1)

	panic(errors.New("出错了"))
}

func _64test2() {

	defer func() {
		fmt.Println(recover())
	}()

	defer func() {
		defer fmt.Println(recover())
		panic(1)
	}()

	defer recover() // 无效

	panic(2)
}
