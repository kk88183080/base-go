package main

import "fmt"

func main() {

	slice := make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2
	_31change(slice...)

	fmt.Println(slice)

	ss := slice[0:2]
	fmt.Printf("ss len:%d, cap:%d\n", len(ss), cap(ss))
	_31change(ss...) // len 2-0=2 cap 5-0=5

	fmt.Println(slice)
}

// ...int 代表切片, 扩容时不会影响原切片里的内容；不扩容时会影响
func _31change(s ...int) {
	fmt.Printf("s len:%d, cap:%d\n", len(s), cap(s))

	s = append(s, 3)
	fmt.Printf("s len:%d, cap:%d\n", len(s), cap(s))
}

//可变函数、append()操作。Go 提供的语法糖...，可以将 slice 传进可变函数，不会创建新的切片。第一次调用 change() 时，
//append() 操作使切片底层数组发生了扩容，原 slice 的底层数组不会改变；第二次调用change() 函数时，使用了操作符[i,j]获得一个新的切片，
//假定为 slice1，它的底层数组和原切片底层数组是重合的，不过 slice1 的长度、容量分别是 2、5，所以在 change() 函数中对 slice1 底层数组的修改会影响到原切片。
