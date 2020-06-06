package main

import "fmt"

func main() {
	//_71test()
	_72test()
}

type _71Slice []int

func New_71Silce() _71Slice {
	return make(_71Slice, 0)
}

func (s *_71Slice) Add(elem int) *_71Slice {
	*s = append(*s, elem)
	fmt.Println(elem)
	return s
}

func _71test() {
	s := New_71Silce()
	defer s.Add(1).Add(2)

	s.Add(3)
}

//1 .Add() 方法的返回值依然是指针类型 *Slice，所以可以循环调用方法 Add()；
//2 .defer 函数的参数（包括接收者）是在 defer 语句出现的位置做计算的，
//而不是在函数执行的时候计算的，所以 s.Add(1) 会先于 s.Add(3) 执行。

func _72test() {
	s := New_71Silce()

	defer func() {
		s.Add(1).Add(2)
	}()

	s.Add(3)
}

// 3 1 2
