package main

import "fmt"

func main() {
	deferCall()
}

/**
defer 执行时有一个压栈，和出栈的过程
*/
func deferCall() {

	defer func() {
		fmt.Println("第一个defer")
	}()

	defer func() {
		fmt.Println("第二个defer")
	}()

	defer func() {
		fmt.Println("第三个defer")
	}()

	panic("我要出错")
}
