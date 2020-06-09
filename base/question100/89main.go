package main

import "fmt"

func main() {
	//_89test()
	_89test1()
}

func _89test() {
	v := []int{1, 2, 3}
	for i, n := 0, len(v); i < n; i++ {
		v = append(v, i)
	}
	fmt.Println(v)
	// {1, 2, 3, 0, 1, 2}
}

type P *int
type Q *int

func _89test1() {
	var p P = new(int)
	*p += 8

	var x *int = p
	var q Q = x
	*q++

	fmt.Println(*p, *q) //9, 9
}
