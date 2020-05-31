package main

import "fmt"

func main() {
	//_24test()
	_24testDefer()
}

func _24test() {
	// map的输出是无序的
	m := map[int]string{0: "zero", 1: "ljw"}

	for k, v := range m {
		fmt.Println(k, v)
	}
}

func _24testDefer() {
	a := 1
	b := 2
	defer _24calc("1", a, _24calc("10", a, b))
	// 1输出 10 1 2 3 --》"1" 1，3
	// 4输出 1  1 3 4

	a = 0
	defer _24calc("2", a, _24calc("20", a, b))
	// 2输出 20 0 2 2 --》"2" 0 2
	// 3输出 2 0 2 2
	b = 1
}

func _24calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
