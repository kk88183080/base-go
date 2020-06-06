package main

import "fmt"

func main() {
	_65test()
}

func _65test() {
	defer func() {
		fmt.Println(recover())
	}()

	defer func() {
		defer func() {
			fmt.Println(recover())
		}()
		panic(1)
	}()

	defer recover() // 无效
	panic(2)
}
