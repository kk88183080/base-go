package main

import (
	"fmt"
	"strings"
)

func main() {
	//_85test()
	//_85test1()
	//_85test2()
	_85test3()
}

func _85test() {
	fmt.Println(strings.TrimRight("ABBA", "BA")) // 输出为空
	fmt.Println(strings.TrimSuffix("ABBA", "BA"))
}

func _85test1() {
	var src, des []int
	src = []int{
		1, 2, 3,
	}
	copy(des, src)
	fmt.Println(des)
}

func _85test2() {
	var src, des []int
	src = []int{1, 2, 3}
	des = make([]int, len(src))

	copy(des, src)
	fmt.Println(des)
}

func _85test3() {
	var src, des []int
	src = []int{1, 3, 56}
	des = append(des, src...)
	fmt.Println(des)
}
