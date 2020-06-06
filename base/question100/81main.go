package main

import "fmt"

func main() {
	_81test()
}

func _81test() {
	var a []int = nil

	a, a[0] = []int{1, 2}, 3
	fmt.Println(a)
}

// 切片数组越界
