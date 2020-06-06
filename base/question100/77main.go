package main

import (
	"fmt"
	"testing"
)

func main() {
	_77test()
}

func _77hello(n ...int) {
	n[0] = 18
}

func _77Test13(t *testing.T) {
	i := []int{5, 6, 7}
	_77hello(i...)
	fmt.Println(i[0]) //18
}

func _77test() {
	t := &testing.T{}
	_77Test13(t)
}

//可变函数是指针传递
