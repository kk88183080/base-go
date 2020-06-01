package main

import (
	"fmt"
	"runtime"
)

func main() {
	//_36testslice()
	_36test1()
}

// 定义函数
func _36f(a, b int) (value int, err error) {
	return 0, nil
}
func _36f1(a int, b int) (value int, err error) {
	return 0, nil
}

// 编译错误
//func _362(a, b int)(value int, error)  {
//	return 0, nil
//}

func _363(a int, b int) (int, int, error) {
	return a, b, nil
}

func _36testslice() {

	// 编译错误，必须传长度
	//s:=make([]int)
	//fmt.Println(s)

	s1 := make([]int, 0)
	fmt.Println(s1)

	s2 := make([]int, 0, 1)
	fmt.Println(s2)
	fmt.Printf("len:%d, cap:%d\n", len(s2), cap(s2))

	s3 := []int{1, 2, 3, 4, 5}
	fmt.Println(s3)
}

/**
  这个方法有可能成功，有可以失败
  select 会随机选择一个可用通道做收发操作，所以可能触发异常，也可能不会
*/
func _36test1() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"

	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}
